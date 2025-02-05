package main

import (
	"bufio"
	"bytes"
	"errors"
	"github.com/mlmon/surveyor/source"
	"os/exec"
	"regexp"
	"strings"
)

func which(binary string) bool {
	err := exec.Command("which", binary).Run()
	if err != nil {
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			return exitError.ExitCode() == 0
		}
	}
	return true
}

var Which = which
var DpkgList = dpkgList

var reDpkg = regexp.MustCompile(`^(\S+)\s+(\S+)\s+(\S+)`)

func Packages() (*source.Records, error) {
	var entries source.Entries
	var hasDpkg = Which("dpkg-query")

	if hasDpkg {
		b, err := DpkgList()
		if err != nil {
			return nil, err
		}
		scanner := bufio.NewScanner(bytes.NewReader(b))
		for scanner.Scan() {
			ln := strings.TrimSpace(scanner.Text())
			a := reDpkg.FindStringSubmatch(ln)
			if len(a) < 4 {
				continue
			}

			// TODO: should probably validate the header and exit early if the output is bad.
			if a[1] != "ii" {
				continue
			}
			entries = append(entries, source.Record{a[2], a[3]})
		}
		if err := scanner.Err(); err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("no package manager found")
	}

	return &source.Records{
		Source:  "package-list",
		Entries: entries,
	}, nil
}

func dpkgList() ([]byte, error) {
	b, err := exec.Command("dpkg-query", "-l").Output()
	if err != nil {
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			return nil, exitError
		}
	}

	return b, nil
}
