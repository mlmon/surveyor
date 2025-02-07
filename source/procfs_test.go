package source_test

import (
	a "github.com/gogunit/gunit/hammy"
	"github.com/mlmon/surveyor/source"
	"testing"
)

func Test_procfs_successful(t *testing.T) {
	assert := a.New(t)
	records, _ := source.ProcFS("testdata/procfs")()
	assert.Is(a.Struct(records).EqualTo(&source.Records{
		Source: "procfs",
		Entries: source.Entries{
			{"fs.file-max", "9223372036854775807"},
			{"kernel.hostname", "nfisher-mbp"},
		}}))
}
