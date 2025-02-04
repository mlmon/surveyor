package main

import (
	"encoding/json"
	"fmt"
	"github.com/mlmon/surveyor/source"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	fns := []source.Fn{
		OsRelease("/etc/os-release"),
		KernelModules("/proc/modules", "/lib/modules"),
		// TODO: Nvidia SMI
		// TODO: lsmod+modinfo
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
		records = append(records, rec)
		logger.Info("processed source", "source", rec.Source, "entries", len(rec.Entries))
	}

	b, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		logger.Error("error marshaling records", "err", err)
	}
	fmt.Printf("%s\n", string(b))
}
