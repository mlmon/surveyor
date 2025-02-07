package source_test

import (
	"errors"
	"github.com/gogunit/gunit/hammy"
	"github.com/mlmon/surveyor/source"
	"os"
	"testing"
)

func Test_packages_failure_when_no_package_manager_found(t *testing.T) {
	defer stubFalseyWhich()()

	assert := hammy.New(t)
	_, err := source.Packages()
	assert.Is(hammy.Error(err))
}

func Test_packages_failure_when_no_package_query_errors(t *testing.T) {
	defer stubTruthyWhich()()
	defer stubDpkgListError()()

	assert := hammy.New(t)
	_, err := source.Packages()
	assert.Is(hammy.Error(err))
}

func Test_debian_packages_success(t *testing.T) {
	defer stubTruthyWhich()()
	defer stubDpkgList()()

	assert := hammy.New(t)

	rec, err := source.Packages()
	assert.Is(hammy.NilError(err))
	assert.Is(hammy.Struct(rec).EqualTo(&source.Records{
		Source: "package-list",
		Entries: source.Entries{
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
	oldDpkgList := source.DpkgList
	source.DpkgList = func() ([]byte, error) {
		return nil, errors.New("permission denied")
	}
	return func() { source.DpkgList = oldDpkgList }
}

func stubDpkgList() func() {
	oldDpkgList := source.DpkgList
	source.DpkgList = func() ([]byte, error) {
		return os.ReadFile("testdata/dpkg-list")
	}
	return func() { source.DpkgList = oldDpkgList }
}
