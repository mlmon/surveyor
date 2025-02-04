package main

import (
	"bufio"
	"github.com/mlmon/surveyor/source"
	"golang.org/x/sys/unix"
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

		entries, err := modInfo(moduleNames, filepath.Join(moduleBasePath, b2s(uname.Release[:])))
		if err != nil {
			return nil, err
		}

		return &source.Records{
			Source:  "kernel-modules",
			Entries: entries,
		}, nil
	}
}

func modInfo(activeModules []string, moduleBasePath string) (source.Entries, error) {
	r, err := os.Open(filepath.Join(moduleBasePath, "modules.dep"))
	if err != nil {
		return nil, err
	}
	defer r.Close()

	m := make(map[string]string)
	_ = m

	var entries source.Entries
	for _, n := range activeModules {
		entries = append(entries, source.Record{Key: n, Value: ""})
	}
	return entries, nil
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
