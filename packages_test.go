package main_test

import (
	"errors"
	a "github.com/gogunit/gunit/hammy"
	"github.com/mlmon/surveyor"
	"github.com/mlmon/surveyor/source"
	"os"
	"testing"
)

func Test_packages_failure_when_no_package_manager_found(t *testing.T) {
	defer stubFalseyWhich()()

	assert := a.New(t)
	_, err := main.Packages()
	assert.Is(a.Error(err))
}

func Test_packages_failure_when_no_package_query_errors(t *testing.T) {
	defer stubTruthyWhich()()
	defer stubDpkgListError()()

	assert := a.New(t)
	_, err := main.Packages()
	assert.Is(a.Error(err))
}

func Test_debian_packages_success(t *testing.T) {
	defer stubTruthyWhich()()
	defer stubDpkgList()()

	assert := a.New(t)

	rec, err := main.Packages()
	assert.Is(a.NilError(err))
	assert.Is(a.Struct(rec).EqualTo(&source.Records{
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

func stubDpkgListError() func() {
	oldDpkgList := main.DpkgList
	main.DpkgList = func() ([]byte, error) {
		return nil, errors.New("permission denied")
	}
	return func() { main.DpkgList = oldDpkgList }
}

func stubDpkgList() func() {
	oldDpkgList := main.DpkgList
	main.DpkgList = func() ([]byte, error) {
		return os.ReadFile("testdata/dpkg-list")
	}
	return func() { main.DpkgList = oldDpkgList }
}
