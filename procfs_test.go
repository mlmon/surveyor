package main_test

import (
	"testing"
)
import . "github.com/mlmon/surveyor"
import "github.com/mlmon/surveyor/source"
import "github.com/gogunit/gunit/hammy"

func Test_procfs_successful(t *testing.T) {
	assert := hammy.New(t)
	records, _ := ProcFS("testdata/procfs")()
	assert.Is(hammy.Struct(records).EqualTo(&source.Records{
		Source: "procfs",
		Entries: []source.Record{
			{"fs.file-max", "9223372036854775807"},
			{"kernel.hostname", "nfisher-mbp"},
		}}))
}
