package main_test

import (
	"github.com/gogunit/gunit/hammy"
	"github.com/mlmon/surveyor"
	"github.com/mlmon/surveyor/source"
	"testing"
)

func Test_os_release_successful(t *testing.T) {
	assert := hammy.New(t)
	rec, _ := main.OsRelease("testdata/os-release")()
	assert.Is(hammy.Struct(rec).EqualTo(&source.Records{
		Source: "os-release",
		Entries: []source.Record{
			{"pretty_name", "Ubuntu 22.04.5 LTS"},
			{"name", "Ubuntu"},
			{"version_id", "22.04"},
			{"version", "22.04.5 LTS (Jammy Jellyfish)"},
			{"version_codename", "jammy"},
		},
	}))
}

func Test_os_release_failure(t *testing.T) {
	assert := hammy.New(t)
	_, err := main.OsRelease("testdata/os-release-missing")()
	assert.Is(hammy.Error(err))
}
