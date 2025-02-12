package main

import (
	"encoding/json"
	"fmt"
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

	path := fmt.Sprintf("bom-%s.cdx.json", sbom.SerialNumber[9:])
	f, err := os.Create(path)
	if err != nil {
		logger.Error("error creating CycloneDX SBOM file", "err", err)
		return 2
	}
	defer f.Close()

	err = json.NewEncoder(f).Encode(sbom)
	if err != nil {
		logger.Error("error writing CycloneDX SBOM", "err", err)
		return 3
	}
	logger.Info("wrote CycloneDX SBOM", "path", path)

	return 0
}

func collect(logger *slog.Logger) *source.RecordSet {
	fns := []source.Fn{
		source.OsRelease("/etc/os-release"),
		source.KernelModules("/proc/modules", "/lib/modules"),
		source.NvidiaSmi,
		source.Packages,
		source.Cmdline("/proc/cmdline"),
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
