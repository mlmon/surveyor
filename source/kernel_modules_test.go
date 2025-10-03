package source_test

import (
	"testing"

	a "github.com/gogunit/gunit/hammy"
	"github.com/mlmon/surveyor/source"
)

func Test_kernel_modules_failure_on_invalid_procPath(t *testing.T) {
	assert := a.New(t)
	_, err := source.KernelModules("testdata/modules-absent", "testdata/modules")()
	assert.Is(a.Error(err))
}

func Test_kernel_modules_success_on_invalid_moduleBasePath(t *testing.T) {
	assert := a.New(t)
	_, err := source.KernelModules("testdata/procfs", "testdata/modules-absent")()
	assert.Is(a.Error(err))
}

func Test_kernel_modules_success(t *testing.T) {
	defer stubUname(unameStub)()

	assert := a.New(t)
	records, err := source.KernelModules("testdata/procfs", "testdata/modules")()
	assert.Is(a.NilError(err))
	assert.Is(a.Struct(records).EqualTo(&source.Records{
		Source: "kernel-modules",
		Entries: source.Entries{
			{Key: "overlay", Value: "6.5.0-1024-aws SMP mod_unload modversions"},
			{Key: "efa", Value: "2.10.0g"},
		},
	}))
}
