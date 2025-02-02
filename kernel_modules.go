package main

import (
	"bufio"
	"github.com/mlmon/surveyor/source"
	"os"
)

func KernelModules(path string) source.Fn {
	return func() (*source.Records, error) {
		r, err := os.Open(path)
		if err != nil {
			return nil, err
		}

		scanner := bufio.NewScanner(r)
		for scanner.Scan() {

		}
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		return &source.Records{
			Source: "kernel-modules",
			Entries: []source.Record{
				{Key: "nvidia", Value: ""},
			},
		}, nil
	}
}
