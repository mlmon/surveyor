package source_test

import (
	a "github.com/gogunit/gunit/hammy"
	"github.com/mlmon/surveyor/source"
	"testing"
)

func Test_cmdline_failure(t *testing.T) {
	assert := a.New(t)
	_, err := source.Cmdline("testdata/invalid-path")()
	assert.Is(a.Error(err))
}

func Test_cmdline_success(t *testing.T) {
	assert := a.New(t)
	records, _ := source.Cmdline("testdata/proc-cmdline")()
	assert.Is(a.Struct(records).EqualTo(&source.Records{
		Source: "cmdline",
		Entries: source.Entries{
			{"BOOT_IMAGE", "/boot/vmlinuz-6.5.0-1024-aws"},
			{"root", "PARTUUID=2a38f7be-dcb6-4780-9e4c-c3537cb2cddd"},
			{"ro", ""},
			{"rd.driver.blacklist", "nouveau"},
			{"nouveau.modeset", "0"},
			{"processor.max_cstate", "1"},
			{"intel_idle.max_cstate", "1"},
			{"console", "tty1"},
			{"console", "ttyS0"},
			{"nvme_core.io_timeout", "4294967295"},
			{"panic", "-1"},
		},
	}))
}
