package source_test

import (
	a "github.com/gogunit/gunit/hammy"
	"github.com/mlmon/surveyor/source"
	"sort"
	"testing"
)

func Test_entries_sort(t *testing.T) {
	assert := a.New(t)

	entries := source.Entries{
		{Key: "key3", Value: "value3"},
		{Key: "key2", Value: "value2"},
		{Key: "key1", Value: "value1"},
	}
	sort.Sort(entries)

	assert.Is(a.Slice(entries).EqualTo(
		source.Record{"key1", "value1"},
		source.Record{"key2", "value2"},
		source.Record{"key3", "value3"},
	))
}
