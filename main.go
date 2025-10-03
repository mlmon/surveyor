package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"sort"

	"github.com/mlmon/surveyor/cyclonedx"
	"github.com/mlmon/surveyor/source"
)

func main() {
	var opts RunOpts

	flag.StringVar(&opts.SbomPath, "out", ".", "Directory to write SBOM to")
	flag.StringVar(&opts.ProcBase, "procbase", "/proc", "Base path for proc fs")
	flag.Parse()

	os.Exit(Run(os.Stdout, &opts))
}

type RunOpts struct {
	SbomPath string
	ProcBase string
}

func (opts *RunOpts) ProcPath(p string) string {
	return filepath.Join(opts.ProcBase, p)
}

func Run(w io.Writer, opts *RunOpts) int {
	logger := slog.New(slog.NewTextHandler(w, nil))

	records := collect(logger, opts)

	sbom, err := cyclonedx.From(records)
	if err != nil {
		logger.Error("error mapping cyclonedx sbom", "err", err)
		return 1
	}

	path := filepath.Join(opts.SbomPath, fmt.Sprintf("bom-%s.cdx.json", sbom.SerialNumber[9:]))
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

func collect(logger *slog.Logger, opts *RunOpts) *source.RecordSet {
	fns := []source.Fn{
		source.OsRelease("/etc/os-release"),
		source.KernelModules(opts.ProcBase, "/lib/modules"),
		source.NvidiaSmi,
		source.Packages,
		source.Cmdline(opts.ProcBase),
		source.ProcSys(opts.ProcBase),
		source.ProcNvidiaParams(opts.ProcBase),
		source.Uname,
	}

	var records []*source.Records
	for _, fn := range fns {
		rec, err := fn()
		if err != nil {
			logger.Error("error processing source", "err", err)
			continue
		}
		// stable sort the entries by key so that 2 or more sets can easily be compared.
		sort.Stable(rec.Entries)
		records = append(records, rec)
		logger.Info("processed source", "source", rec.Source, "entries", len(rec.Entries))
	}

	sort.SliceStable(records, func(i, j int) bool {
		return records[i].Source < records[j].Source
	})
	return &source.RecordSet{
		Records: records,
	}
}
