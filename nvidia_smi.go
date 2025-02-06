package main

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/mlmon/surveyor/source"
	"os/exec"
	"strings"
)

var NvidiaQuery = nvidiaQuery

// nvidia-smi --format=csv --query-gpu=gpu_name,vbios_version,driver_version,inforom.oem,inforom.ecc,inforom.img,compute_cap
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

	if len(rows) < 2 || len(rows[0]) != len(rows[1]) {
		return nil, fmt.Errorf("expected %d columns, got %d", len(rows[0]), len(rows[1]))
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
	b, err := exec.Command("nvidia-smi", "--format=csv", "--query-gpu=gpu_name,vbios_version,driver_version,inforom.oem,inforom.ecc,inforom.img,compute_cap").Output()
	if err != nil {
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			return nil, exitError
		}
	}

	return b, nil
}
