package cyclonedx_test

import (
	a "github.com/gogunit/gunit/hammy"
	"github.com/mlmon/surveyor/cyclonedx"
	"github.com/mlmon/surveyor/source"

	"testing"
)

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

	assert.Is(a.Struct(sbom).EqualTo(&cyclonedx.SBOM{
		BomFormat:   "From",
		SpecVersion: "1.4",
		Version:     "1",
		Components: []cyclonedx.Component{
			{Name: "adduser", Version: "3.118ubuntu5", Type: "library"},
		},
	}))
}
