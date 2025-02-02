package main_test

import (
	"github.com/gogunit/gunit/hammy"
	"github.com/mlmon/surveyor"
	"github.com/mlmon/surveyor/source"
	"os"
	"testing"
)

func Test_which_fails_with_non_existent_binary(t *testing.T) {
	assert := hammy.New(t)
	hasBinary := main.Which("plants-and-bees-knees")
	assert.Is(hammy.False(hasBinary))
}

func Test_which_succeeds_with_sh_binary(t *testing.T) {
	assert := hammy.New(t)
	hasBinary := main.Which("sh")
	assert.Is(hammy.True(hasBinary))
}

func Test_packages_fails_when_no_package_manager_found(t *testing.T) {
	defer stubFalseyWhich()()

	assert := hammy.New(t)
	_, err := main.Packages()
	assert.Is(hammy.Error(err))
}

func Test_debian_packages_success(t *testing.T) {
	defer stubTruthyWhich()()
	defer stubDpkgList()()

	assert := hammy.New(t)

	rec, err := main.Packages()
	assert.Is(hammy.NilError(err))
	assert.Is(hammy.Struct(rec).EqualTo(&source.Records{
		Source: "package-list",
		Entries: []source.Record{
			{Key: "accountsservice", Value: "22.07.5-2ubuntu1.5"},
			{Key: "acl", Value: "2.3.1-1"},
			{Key: "acpi-support", Value: "0.144"},
			{Key: "acpid", Value: "1:2.0.33-1ubuntu1"},
			{Key: "adduser", Value: "3.118ubuntu5"},
			{Key: "adwaita-icon-theme", Value: "41.0-1ubuntu1"},
		},
	}))
}

func stubDpkgList() func() {
	oldDpkgList := main.DpkgList
	main.DpkgList = func() ([]byte, error) {
		return os.ReadFile("testdata/dpkg-list")
	}
	return func() { main.DpkgList = oldDpkgList }
}

func stubFalseyWhich() func() {
	oldWhich := main.Which
	main.Which = func(binary string) bool { return false }
	return func() { main.Which = oldWhich }
}

func stubTruthyWhich() func() {
	oldWhich := main.Which
	main.Which = func(binary string) bool {
		return true
	}
	return func() { main.Which = oldWhich }
}
