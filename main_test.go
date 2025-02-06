//go:build linux

package main_test

import (
	"bytes"
	a "github.com/gogunit/gunit/hammy"
	"github.com/mlmon/surveyor"
	"testing"
)

func Test_run(t *testing.T) {
	assert := a.New(t)

	var buf bytes.Buffer
	main.Run(&buf)

	s := buf.String()
	assert.Is(a.String(s).Contains(`level=INFO msg="processed source" source=os-release entries=`))
	assert.Is(a.String(s).Contains(`level=INFO msg="processed source" source=kernel-modules entries=`))
	assert.Is(a.String(s).Contains(`level=INFO msg="processed source" source=package-list entries=`))
	assert.Is(a.String(s).Contains(`level=INFO msg="processed source" source=procfs entries=`))
	assert.Is(a.String(s).Contains(`level=INFO msg="processed source" source=uname entries=`))
}
