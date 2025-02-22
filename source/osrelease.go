package source

import (
	"bufio"
	"os"
	"strings"
)

func OsRelease(path string) Fn {
	var accept = map[string]bool{
		"pretty_name":      true,
		"name":             true,
		"version_id":       true,
		"version":          true,
		"version_codename": true,
	}
	return func() (*Records, error) {
		r, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer r.Close()

		var entries []Record
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			ln := scanner.Text()
			ln = strings.TrimSpace(ln)

			// split string into key and value
			a := strings.Split(ln, "=")
			if len(a) != 2 {
				continue
			}

			k := strings.ToLower(strings.TrimSpace(a[0]))
			v := strings.TrimSpace(a[1])

			// only retrieve keys of interest.
			if !accept[k] {
				continue
			}

			// remove double quotes from value
			if v[0] == '"' && v[len(v)-1] == '"' {
				v = v[1 : len(v)-1]
			}
			entries = append(entries, Record{Key: k, Value: v})
		}
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		return &Records{
			Source:  "os-release",
			Entries: entries,
		}, nil
	}
}
