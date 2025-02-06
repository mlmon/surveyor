package main

import (
	"bytes"
	"encoding/csv"
	"errors"
	"github.com/mlmon/surveyor/source"
	"os/exec"
	"strings"
)

var NvidiaQuery = nvidiaQuery

func NvidiaSmi() (*source.Records, error) {
	var entries source.Entries
	var hasNvidiaSmi = Which("nvidia-smi")
	if !hasNvidiaSmi {
		return nil, errors.New("nvidia-smi not found")
	}

	b, err := NvidiaQuery()
	if err != nil {
		return nil, err
	}

	rows, err := csv.NewReader(bytes.NewReader(b)).ReadAll()
	if err != nil {
		return nil, err
	}

	if len(rows) < 2 {
		return nil, errors.New("expected at least 1 data row, got header")
	}

	for i, name := range rows[0] {
		entries = append(entries, source.Record{Key: strings.TrimSpace(name), Value: strings.TrimSpace(rows[1][i])})
	}

	return &source.Records{
		Source:  "nvidia-smi",
		Entries: entries,
	}, nil
}

func nvidiaQuery() ([]byte, error) {
	return exec.Command("nvidia-smi", "--format=csv", "--query-gpu=gpu_name,vbios_version,driver_version,inforom.oem,inforom.ecc,inforom.img,compute_cap").Output()
}
