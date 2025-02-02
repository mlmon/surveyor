package main

import (
	"bytes"
	"github.com/mlmon/surveyor/source"
	"golang.org/x/sys/unix"
)

var UnixUname = unix.Uname

func Uname() (*source.Records, error) {
	var uname unix.Utsname
	err := UnixUname(&uname)
	if err != nil {
		return nil, err
	}
	var records []source.Record

	records = append(records, source.Record{Key: "machine", Value: b2s(uname.Machine[:])})
	records = append(records, source.Record{Key: "nodename", Value: b2s(uname.Nodename[:])})
	records = append(records, source.Record{Key: "release", Value: b2s(uname.Release[:])})
	records = append(records, source.Record{Key: "sysname", Value: b2s(uname.Sysname[:])})
	records = append(records, source.Record{Key: "version", Value: b2s(uname.Version[:])})

	return &source.Records{
		Source:  "uname",
		Entries: records,
	}, nil
}

// b2s converts a null terminated byte array to a string.
func b2s(b []byte) string {
	n := bytes.IndexByte(b, 0)
	return string(b[:n])
}
