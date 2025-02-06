package main

import (
	"encoding/json"
	"fmt"
	"github.com/mlmon/surveyor/source"
	"log/slog"
	"os"
	"sort"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	fns := []source.Fn{
		OsRelease("/etc/os-release"),
		KernelModules("/proc/modules", "/lib/modules"),
		NvidiaSmi,
		Packages,
		ProcFS("/proc/sys"),
		Uname,
	}

	var records []*source.Records
	for _, fn := range fns {
		rec, err := fn()
		if err != nil {
			logger.Error("error processing source", "err", err)
			continue
		}
		// sort the entries by key so that 2 or more sets can easily be compared.
		sort.Sort(rec.Entries)
		records = append(records, rec)
		logger.Info("processed source", "source", rec.Source, "entries", len(rec.Entries))
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].Source < records[j].Source
	})

	b, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		logger.Error("error marshaling records", "err", err)
	}
	fmt.Printf("%s\n", string(b))
}
