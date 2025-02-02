package main_test

import (
	a "github.com/gogunit/gunit/hammy"
	"github.com/mlmon/surveyor"
	"github.com/mlmon/surveyor/source"
	"testing"
)

func Test_kernel_modules_failure(t *testing.T) {
	assert := a.New(t)
	_, err := main.KernelModules("testdata/proc-modules-absent")()
	assert.Is(a.Error(err))
}

func Test_kernel_modules_success(t *testing.T) {
	assert := a.New(t)
	records, _ := main.KernelModules("testdata/proc-modules")()
	assert.Is(a.Struct(records).EqualTo(&source.Records{
		Source: "kernel-modules",
	}))
}
