package source_test

import (
	a "github.com/gogunit/gunit/hammy"
	"github.com/mlmon/surveyor/source"
	"testing"
)

func Test_os_release_successful(t *testing.T) {
	assert := a.New(t)
	rec, _ := source.OsRelease("testdata/os-release")()
	assert.Is(a.Struct(rec).EqualTo(&source.Records{
		Source: "os-release",
		Entries: source.Entries{
			{"pretty_name", "Ubuntu 22.04.5 LTS"},
			{"name", "Ubuntu"},
			{"version_id", "22.04"},
			{"version", "22.04.5 LTS (Jammy Jellyfish)"},
			{"version_codename", "jammy"},
		},
	}))
}

func Test_os_release_failure(t *testing.T) {
	assert := a.New(t)
	_, err := source.OsRelease("testdata/os-release-missing")()
	assert.Is(a.Error(err))
}
