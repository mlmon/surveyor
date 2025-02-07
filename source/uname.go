package source

import (
	"bytes"
	"golang.org/x/sys/unix"
)

var UnixUname = unix.Uname

func Uname() (*Records, error) {
	var uname unix.Utsname
	err := UnixUname(&uname)
	if err != nil {
		return nil, err
	}
	var records []Record

	records = append(records, Record{Key: "machine", Value: b2s(uname.Machine[:])})
	records = append(records, Record{Key: "nodename", Value: b2s(uname.Nodename[:])})
	records = append(records, Record{Key: "release", Value: b2s(uname.Release[:])})
	records = append(records, Record{Key: "sysname", Value: b2s(uname.Sysname[:])})
	records = append(records, Record{Key: "version", Value: b2s(uname.Version[:])})

	return &Records{
		Source:  "uname",
		Entries: records,
	}, nil
}

// b2s converts a null terminated byte array to a string.
func b2s(b []byte) string {
	n := bytes.IndexByte(b, 0)
	return string(b[:n])
}
