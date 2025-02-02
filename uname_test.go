package main_test

import (
	"errors"
	"github.com/gogunit/gunit/hammy"
	"github.com/mlmon/surveyor"
	"github.com/mlmon/surveyor/source"
	"golang.org/x/sys/unix"
	"testing"
)

func Test_uname_successful(t *testing.T) {
	defer stubUname(unameStub)()

	assert := hammy.New(t)
	records, _ := main.Uname()
	assert.Is(hammy.Struct(records).EqualTo(&source.Records{
		Source: "uname",
		Entries: []source.Record{
			{Key: "machine", Value: "aarch64"},
			{Key: "nodename", Value: "715bf308c176"},
			{Key: "release", Value: "6.12.5-orbstack-00287-gf8da5d508983"},
			{Key: "sysname", Value: "Linux"},
			{Key: "version", Value: "#19 SMP Tue Dec 17 08:07:20 UTC 2024"},
		},
	}))
}

func Test_uname_failure(t *testing.T) {
	defer stubUname(unameErrorStub)()

	assert := hammy.New(t)
	_, err := main.Uname()
	assert.Is(hammy.Error(err))
}

func stubUname(fn func(*unix.Utsname) error) func() {
	old := main.UnixUname
	main.UnixUname = fn
	return func() { main.UnixUname = old }
}

func unameStub(uname *unix.Utsname) error {
	copy(uname.Machine[:], "aarch64")
	copy(uname.Nodename[:], "715bf308c176")
	copy(uname.Release[:], "6.12.5-orbstack-00287-gf8da5d508983")
	copy(uname.Sysname[:], "Linux")
	copy(uname.Version[:], "#19 SMP Tue Dec 17 08:07:20 UTC 2024")
	return nil
}

func unameErrorStub(uname *unix.Utsname) error {
	return errors.New("fake error")
}
