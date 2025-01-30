package main

import (
	"bytes"
	"golang.org/x/sys/unix"
)

func Uname() (*SourceRecords, error) {
	var uname unix.Utsname
	err := unix.Uname(&uname)
	if err != nil {
		return nil, err
	}
	var records []Record

	records = append(records, Record{"machine", b2s(uname.Machine[:])})
	records = append(records, Record{"nodename", b2s(uname.Nodename[:])})
	records = append(records, Record{"release", b2s(uname.Release[:])})
	records = append(records, Record{"sysname", b2s(uname.Sysname[:])})
	records = append(records, Record{"version", b2s(uname.Version[:])})
	return &SourceRecords{
		Source:  "uname",
		Records: records,
	}, nil
}

// b2s converts a null terminated byte array to a string.
func b2s(b []byte) string {
	n := bytes.IndexByte(b, 0)
	return string(b[:n])
}
