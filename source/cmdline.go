package source

import (
	"os"
	"strings"
)

func Cmdline(path string) Fn {
	return func() (*Records, error) {
		b, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		s := string(b)
		a := strings.Split(s, " ")
		var entries Entries
		for _, e := range a {
			if strings.Contains(e, "=") {
				kv := strings.SplitN(e, "=", 2)
				entries = append(entries, Record{kv[0], kv[1]})
				continue
			}
			entries = append(entries, Record{e, ""})
		}
		return &Records{
			Source:  "cmdline",
			Entries: entries,
		}, nil
	}
}
