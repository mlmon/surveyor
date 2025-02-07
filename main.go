package main

import (
	"github.com/mlmon/surveyor/cyclonedx"
	"github.com/mlmon/surveyor/source"
	"io"
	"log/slog"
	"os"
	"sort"
)

func main() {
	os.Exit(Run(os.Stdout))
}

func Run(w io.Writer) int {
	logger := slog.New(slog.NewTextHandler(w, nil))

	records := collect(logger)

	sbom, err := cyclonedx.From(records)
	if err != nil {
		logger.Error("error mapping cyclonedx sbom", "err", err)
		return 1
	}
	_ = sbom

	return 0
}

func collect(logger *slog.Logger) *source.RecordSet {
	fns := []source.Fn{
		source.OsRelease("/etc/os-release"),
		source.KernelModules("/proc/modules", "/lib/modules"),
		source.NvidiaSmi,
		source.Packages,
		source.ProcFS("/proc/sys"),
		source.Uname,
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
	return &source.RecordSet{
		Records: records,
	}
}
