package source

import (
	"os"
	"path/filepath"
	"strings"
)

const Procfs = "procfs"
const NvidiaDriverParams = "nvidia_driver_params"

func ProcSys(path string) Fn {
	procfs := filepath.Join(path, "sys")
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
			Source:  Procfs,
			Entries: records,
		}, err
	}
}

func ProcNvidiaParams(path string) Fn {
	params := filepath.Join(path, "driver/nvidia/params")
	return func() (*Records, error) {
		var records []Record

		data, err := os.ReadFile(params)
		if err != nil {
			return nil, err
		}

		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}
			
			parts := strings.SplitN(line, ": ", 2)
			if len(parts) == 2 {
				record := Record{
					Key:   parts[0],
					Value: parts[1],
				}
				records = append(records, record)
			}
		}

		return &Records{
			Source:  NvidiaDriverParams,
			Entries: records,
		}, nil
	}
}
