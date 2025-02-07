package source

import (
	"bufio"
	"bytes"
	"errors"
	"os/exec"
	"regexp"
	"strings"
)

var DpkgList = dpkgList

var reDpkg = regexp.MustCompile(`^(\S+)\s+(\S+)\s+(\S+)`)

const PackageList = "package-list"

func Packages() (*Records, error) {
	var entries Entries
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
			entries = append(entries, Record{Key: a[2], Value: a[3]})
		}
		if err := scanner.Err(); err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("no package manager found")
	}

	return &Records{
		Source:  PackageList,
		Entries: entries,
	}, nil
}

func dpkgList() ([]byte, error) {
	return exec.Command("dpkg-query", "-l").Output()
}
