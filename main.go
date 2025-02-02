package main

import (
	"github.com/mlmon/surveyor/source"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	fns := []source.Fn{
		OsRelease("/etc/os-release"),
		KernelModules("/proc/modules"),
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
}
