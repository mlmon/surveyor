package main

import (
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("hello world")
}

type SourceFn func() (*SourceRecords, error)

type SourceRecords struct {
	Source  string   `json:"source"`
	Records []Record `json:"records"`
}

type Record struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
