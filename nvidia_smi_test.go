package main_test

import (
	"errors"
	a "github.com/gogunit/gunit/hammy"
	"github.com/mlmon/surveyor"
	"github.com/mlmon/surveyor/source"
	"testing"
)

func Test_nvidia_smi_failure_binary_not_found(t *testing.T) {
	defer stubFalseyWhich()()

	assert := a.New(t)
	_, err := main.NvidiaSmi()
	assert.Is(a.Error(err))
}

func Test_nvidia_smi_failure_command_error(t *testing.T) {
	defer stubTruthyWhich()()
	defer stubNvidiaQueryError()()

	assert := a.New(t)
	_, err := main.NvidiaSmi()
	assert.Is(a.Error(err))
}

func Test_nvidia_smi_failure_truncated_output(t *testing.T) {
	defer stubTruthyWhich()()
	defer stubNvidiaQueryTruncated()()

	assert := a.New(t)
	_, err := main.NvidiaSmi()
	assert.Is(a.Error(err))
}

func Test_nvidia_smi_failure_no_table_body(t *testing.T) {
	defer stubTruthyWhich()()
	defer stubNvidiaQueryHeaderOnly()()

	assert := a.New(t)
	_, err := main.NvidiaSmi()
	assert.Is(a.Error(err))
}

func Test_nvidia_smi_success(t *testing.T) {
	defer stubTruthyWhich()()
	defer stubNvidiaQuery()()

	assert := a.New(t)
	rec, _ := main.NvidiaSmi()
	assert.Is(a.Struct(rec).EqualTo(&source.Records{
		Source: "nvidia-smi",
		Entries: []source.Record{
			{"name", "NVIDIA H100 80GB HBM3"},
			{"vbios_version", "96.00.A5.00.01"},
			{"driver_version", "550.90.07"},
			{"inforom.oem", "2.1"},
			{"inforom.ecc", "7.16"},
			{"inforom.img", "G520.0200.00.05"},
			{"compute_cap", "9.0"},
		},
	}))
}

func stubNvidiaQueryError() func() {
	oldQuery := main.NvidiaQuery
	main.NvidiaQuery = func() ([]byte, error) {
		return nil, errors.New("stub error")
	}
	return func() { main.NvidiaQuery = oldQuery }
}

func stubNvidiaQueryTruncated() func() {
	oldQuery := main.NvidiaQuery
	main.NvidiaQuery = func() ([]byte, error) {
		return []byte(`name, vbios_version, driver_version, inforom.oem, inforom.ecc, inforom.img, compute_cap
NVIDIA H100 80GB HBM3, 96.00.`), nil
	}
	return func() { main.NvidiaQuery = oldQuery }
}

func stubNvidiaQuery() func() {
	oldQuery := main.NvidiaQuery
	main.NvidiaQuery = func() ([]byte, error) {
		return []byte(`name, vbios_version, driver_version, inforom.oem, inforom.ecc, inforom.img, compute_cap
NVIDIA H100 80GB HBM3, 96.00.A5.00.01, 550.90.07, 2.1, 7.16, G520.0200.00.05, 9.0
NVIDIA H100 80GB HBM3, 96.00.A5.00.01, 550.90.07, 2.1, 7.16, G520.0200.00.05, 9.0
NVIDIA H100 80GB HBM3, 96.00.A5.00.01, 550.90.07, 2.1, 7.16, G520.0200.00.05, 9.0
NVIDIA H100 80GB HBM3, 96.00.A5.00.01, 550.90.07, 2.1, 7.16, G520.0200.00.05, 9.0
NVIDIA H100 80GB HBM3, 96.00.A5.00.01, 550.90.07, 2.1, 7.16, G520.0200.00.05, 9.0
NVIDIA H100 80GB HBM3, 96.00.A5.00.01, 550.90.07, 2.1, 7.16, G520.0200.00.05, 9.0
NVIDIA H100 80GB HBM3, 96.00.A5.00.01, 550.90.07, 2.1, 7.16, G520.0200.00.05, 9.0
NVIDIA H100 80GB HBM3, 96.00.A5.00.01, 550.90.07, 2.1, 7.16, G520.0200.00.05, 9.0`), nil
	}
	return func() { main.NvidiaQuery = oldQuery }
}

func stubNvidiaQueryHeaderOnly() func() {
	oldQuery := main.NvidiaQuery
	main.NvidiaQuery = func() ([]byte, error) {
		return []byte(`name, vbios_version, driver_version, inforom.oem, inforom.ecc, inforom.img, compute_cap`), nil
	}
	return func() { main.NvidiaQuery = oldQuery }
}
