package main

import (
	"os"
	"path/filepath"
	"strings"
)

func ProcFS(procfs string) SourceFn {
	return func() (*SourceRecords, error) {
		var records []Record

		err := filepath.Walk(procfs, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				// If we have a permissions error or similar, just skip this entry
				return nil
			}

			// skip directories
			if info.IsDir() {
				return nil
			}

			// Read the file
			data, err := os.ReadFile(path)
			if err != nil {
				// Could not read file (permission error, etc.), skip
				return nil
			}
			value := strings.TrimSpace(string(data))

			// Convert the path `/proc/sys/net/ipv4/ip_forward` to `net.ipv4.ip_forward`
			// 1. Strip off "/proc/sys/"
			relative := strings.TrimPrefix(path, procfs+"/")
			// 2. Replace "/" with "."
			paramName := strings.ReplaceAll(relative, "/", ".")

			records = append(records, Record{Key: paramName, Value: value})

			return nil
		})

		if err != nil {
			return nil, err
		}

		return &SourceRecords{
			Source:  "procfs",
			Records: records,
		}, nil
	}
}
