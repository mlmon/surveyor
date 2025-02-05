package main

import (
	"bufio"
	"debug/elf"
	"errors"
	"github.com/mlmon/surveyor/source"
	"golang.org/x/sys/unix"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func KernelModules(procPath string, moduleBasePath string) source.Fn {
	return func() (*source.Records, error) {
		moduleNames, err := readProcModules(procPath)
		if err != nil {
			return nil, err
		}

		var uname unix.Utsname
		err = UnixUname(&uname)
		if err != nil {
			return nil, err
		}

		entries, err := moduleVersions(moduleNames, filepath.Join(moduleBasePath, b2s(uname.Release[:])))
		if err != nil {
			return nil, err
		}

		return &source.Records{
			Source:  "kernel-modules",
			Entries: entries,
		}, nil
	}
}

func moduleVersions(activeModules []string, moduleBasePath string) (source.Entries, error) {
	r, err := os.Open(filepath.Join(moduleBasePath, "modules.dep"))
	if err != nil {
		return nil, err
	}
	defer r.Close()

	m := make(map[string]string)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}
		p := filepath.Join(moduleBasePath, parts[0])
		meta, err := modInfo(p)
		if err != nil {
			// TODO: replace with logger from main or make error collection that is returned.
			log.Printf("error parsing module %q: %v", p, err)
			continue
		}
		m[meta.Name] = meta.Version
	}

	var entries source.Entries
	for _, n := range activeModules {
		entries = append(entries, source.Record{Key: n, Value: m[n]})
	}
	return entries, nil
}

type ModuleMetadata struct {
	Name    string
	Version string
}

func modInfo(modPath string) (*ModuleMetadata, error) {
	ef, err := elf.Open(modPath)
	if err != nil {
		return nil, err
	}
	defer ef.Close()

	sec := ef.Section(".modinfo")
	if sec == nil {
		return nil, errors.New("no .modinfo section found")
	}

	data, err := sec.Data()
	if err != nil {
		return nil, err
	}

	var metadata ModuleMetadata
	var vermagic string
	entries := strings.Split(string(data), "\x00")
	for _, entry := range entries {
		if strings.HasPrefix(entry, "name=") {
			metadata.Name = strings.TrimPrefix(entry, "name=")
		} else if strings.HasPrefix(entry, "version=") {
			metadata.Version = strings.TrimPrefix(entry, "version=")
		} else if strings.HasPrefix(entry, "vermagic=") {
			vermagic = strings.TrimSpace(strings.TrimPrefix(entry, "vermagic="))
		}
	}

	if metadata.Name == "" {
		return nil, errors.New("module name not found in .modinfo")
	}

	// version field is typically missing or empty for modules shipped with the kernel.
	// if empty use vermagic which is the kernel generated configuration identifier.
	if metadata.Version == "" {
		metadata.Version = vermagic
	}

	return &metadata, nil
}

func readProcModules(procPath string) ([]string, error) {
	// read from proc modules
	r, err := os.Open(procPath)
	if err != nil {
		return nil, err
	}

	var modules []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		ln := strings.TrimSpace(scanner.Text())
		a := strings.SplitN(ln, " ", 2)
		if len(a) != 2 {
			continue
		}

		module := strings.TrimSpace(a[0])

		modules = append(modules, module)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return modules, nil
}
