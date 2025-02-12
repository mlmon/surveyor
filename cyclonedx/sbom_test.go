package cyclonedx_test

import (
	a "github.com/gogunit/gunit/hammy"
	"github.com/mlmon/surveyor/cyclonedx"
	"github.com/mlmon/surveyor/source"

	"testing"
)

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

func Test_From_empty_recordset_yields_only_schema_headers(t *testing.T) {
	defer stubRandom()()

	assert := a.New(t)
	records := source.RecordSet{Records: []*source.Records{}}
	sbom, _ := cyclonedx.From(&records)
	assert.Is(a.Struct(sbom).EqualTo(&cyclonedx.SBOM{
		BomFormat:    "CycloneDX",
		SpecVersion:  "1.4",
		Version:      "1",
		SerialNumber: "urn:uuid:ffffffff-ffff-ffff-ffff-ffffffffffff",
	}))
}

func Test_From_maps_cmdline_to_component(t *testing.T) {
	assert := a.New(t)
	records := source.RecordSet{
		Records: []*source.Records{
			{
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
			},
		},
	}

	sbom, _ := cyclonedx.From(&records)

	assert.Is(SBOM(sbom).FirstComponent(
		cyclonedx.Component{
			Tags:    []string{"cmdline"},
			Name:    "BOOT_IMAGE",
			Version: "/boot/vmlinuz-6.5.0-1024-aws",
			Type:    cyclonedx.File}))
}

func Test_From_maps_procfs_to_component(t *testing.T) {
	assert := a.New(t)
	records := source.RecordSet{
		Records: []*source.Records{
			{
				Source: source.Procfs,
				Entries: source.Entries{
					{Key: "fs.file-max", Value: "9223372036854775807"},
				},
			},
		},
	}

	sbom, _ := cyclonedx.From(&records)

	assert.Is(SBOM(sbom).FirstComponent(
		cyclonedx.Component{
			Name:    "fs.file-max",
			Version: "9223372036854775807",
			Type:    cyclonedx.File,
			Tags:    []string{source.Procfs}}))
}

func Test_From_maps_uname_to_component(t *testing.T) {
	assert := a.New(t)
	records := source.RecordSet{
		Records: []*source.Records{
			{
				Source: "uname",
				Entries: source.Entries{
					{Key: "release", Value: "6.5.0-1024-aws"},
				},
			},
		},
	}

	sbom, _ := cyclonedx.From(&records)
	assert.Is(SBOM(sbom).FirstComponent(
		cyclonedx.Component{
			Name:    "release",
			Version: "6.5.0-1024-aws",
			Type:    cyclonedx.OperatingSystem,
			Tags:    []string{"uname"},
		}))
}

func Test_From_maps_nvidia_smi_firmware_to_component(t *testing.T) {
	assert := a.New(t)
	records := source.RecordSet{
		Records: []*source.Records{
			{
				Source: "nvidia-smi",
				Entries: source.Entries{
					{Key: "vbios_version", Value: "96.00.A5.00.01"},
				},
			},
		},
	}

	sbom, _ := cyclonedx.From(&records)
	assert.Is(SBOM(sbom).FirstComponent(
		cyclonedx.Component{
			Name:    "vbios",
			Version: "96.00.A5.00.01",
			Type:    cyclonedx.Firmware,
			Tags:    []string{"nvidia-smi"},
		}))
}

func Test_From_maps_nvidia_smi_gpu_to_component(t *testing.T) {
	assert := a.New(t)
	records := source.RecordSet{
		Records: []*source.Records{
			{
				Source: "nvidia-smi",
				Entries: source.Entries{
					{Key: "name", Value: "NVIDIA H100 80GB HBM3"},
				},
			},
		},
	}

	sbom, _ := cyclonedx.From(&records)
	assert.Is(SBOM(sbom).FirstComponent(
		cyclonedx.Component{
			Name:    "gpu",
			Version: "NVIDIA H100 80GB HBM3",
			Type:    cyclonedx.Device,
			Tags:    []string{"nvidia-smi"},
		}))
}

func Test_From_maps_package_list_to_component(t *testing.T) {
	assert := a.New(t)
	records := source.RecordSet{
		Records: []*source.Records{
			{
				Source: source.PackageList,
				Entries: source.Entries{
					{Key: "adduser", Value: "3.118ubuntu5"},
				},
			},
		},
	}

	sbom, _ := cyclonedx.From(&records)
	assert.Is(SBOM(sbom).FirstComponent(
		cyclonedx.Component{
			Name:    "adduser",
			Version: "3.118ubuntu5",
			Type:    cyclonedx.Library,
			Tags:    []string{source.PackageList},
		}))
}

func SBOM(model *cyclonedx.SBOM) *sbom {
	return &sbom{model}
}

type sbom struct {
	model *cyclonedx.SBOM
}

func (s *sbom) FirstComponent(component cyclonedx.Component) a.AssertionMessage {
	actual := s.model.Components[0]
	return a.Struct(actual).EqualTo(component)
}
