package source_test

import (
	"errors"
	a "github.com/gogunit/gunit/hammy"
	"github.com/mlmon/surveyor/source"
	"golang.org/x/sys/unix"
	"testing"
)

func Test_uname_successful(t *testing.T) {
	defer stubUname(unameStub)()

	assert := a.New(t)
	records, _ := source.Uname()
	assert.Is(a.Struct(records).EqualTo(&source.Records{
		Source: "uname",
		Entries: source.Entries{
			{Key: "machine", Value: "aarch64"},
			{Key: "nodename", Value: "715bf308c176"},
			{Key: "release", Value: "6.5.0-1024-aws"},
			{Key: "sysname", Value: "Linux"},
			{Key: "version", Value: "#19 SMP Tue Dec 17 08:07:20 UTC 2024"},
		},
	}))
}

func Test_uname_failure(t *testing.T) {
	defer stubUname(unameErrorStub)()

	assert := a.New(t)
	_, err := source.Uname()
	assert.Is(a.Error(err))
}

func stubUname(fn func(*unix.Utsname) error) func() {
	old := source.UnixUname
	source.UnixUname = fn
	return func() { source.UnixUname = old }
}

func unameStub(uname *unix.Utsname) error {
	copy(uname.Machine[:], "aarch64")
	copy(uname.Nodename[:], "715bf308c176")
	copy(uname.Release[:], "6.5.0-1024-aws")
	copy(uname.Sysname[:], "Linux")
	copy(uname.Version[:], "#19 SMP Tue Dec 17 08:07:20 UTC 2024")
	return nil
}

func unameErrorStub(uname *unix.Utsname) error {
	return errors.New("fake error")
}
