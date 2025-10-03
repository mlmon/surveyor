//go:build linux

package main_test

import (
	"bytes"
	"testing"

	a "github.com/gogunit/gunit/hammy"
	"github.com/mlmon/surveyor"
	"github.com/mlmon/surveyor/cyclonedx"
)

func Test_run(t *testing.T) {
	defer stubRandom()()

	assert := a.New(t)

	var buf bytes.Buffer
	rc := main.Run(&buf, main.RunOpts{SbomPath: "."})

	assert.Is(a.Number(rc).IsZero())

	s := buf.String()
	assert.Is(a.String(s).Contains(`level=INFO msg="processed source" source=os-release entries=`))
	assert.Is(a.String(s).Contains(`level=INFO msg="processed source" source=kernel-modules entries=`))
	assert.Is(a.String(s).Contains(`level=INFO msg="processed source" source=package-list entries=`))
	assert.Is(a.String(s).Contains(`level=INFO msg="processed source" source=procfs entries=`))
	assert.Is(a.String(s).Contains(`level=INFO msg="processed source" source=uname entries=`))
}

func stubRandom() func() {
	o := cyclonedx.Random
	cyclonedx.Random = func(b []byte) (n int, err error) {
		for i := range b {
			b[i] = 0xff
		}
		return len(b), nil
	}
	return func() {
		cyclonedx.Random = o
	}
}
