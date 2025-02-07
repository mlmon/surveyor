package source

import (
	"os"
	"path/filepath"
	"strings"
)

func ProcFS(procfs string) Fn {
	return func() (*Records, error) {
		var records []Record

		err := filepath.Walk(procfs, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}

			// Skip directories
			if info.IsDir() {
				return nil
			}

			data, err := os.ReadFile(path)
			if err != nil {
				// File unreadable (permission error, etc.), skip
				return nil
			}
			value := strings.TrimSpace(string(data))

			// Convert from path `/proc/sys/net/ipv4/ip_forward` to `net.ipv4.ip_forward`
			// Trim "/proc/sys/"
			relative := strings.TrimPrefix(path, procfs+"/")
			// Replace "/" with "."
			paramName := strings.ReplaceAll(relative, "/", ".")

			records = append(records, Record{Key: paramName, Value: value})

			return nil
		})

		return &Records{
			Source:  "procfs",
			Entries: records,
		}, err
	}
}
