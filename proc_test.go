package main_test

import "testing"

import . "github.com/mlmon/surveyor"

import "github.com/gogunit/gunit/hammy"

func Test_standard_read_successful(t *testing.T) {
	assert := hammy.New(t)
	records, _ := ProcFS("testdata/procfs")
	assert.Is(hammy.Struct(records).EqualTo(&SourceRecords{
		Source: "procfs",
		Records: []Record{
			{"fs.file-max", "9223372036854775807"},
			{"kernel.hostname", "nfisher-mbp"},
		},
	}))
}
