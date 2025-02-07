package cyclonedx_test

import (
	a "github.com/gogunit/gunit/hammy"
	"github.com/mlmon/surveyor/cyclonedx"
	"testing"
)

func Test_uuid_without_error(t *testing.T) {
	assert := a.New(t)
	_, err := cyclonedx.Uuid()
	assert.Is(a.NilError(err))
}

func Test_uuid_returns_uuid(t *testing.T) {
	assert := a.New(t)
	uid, _ := cyclonedx.Uuid()
	assert.IsNot(a.String(uid).EqualTo("00000000-0000-0000-0000-000000000000"))
}
