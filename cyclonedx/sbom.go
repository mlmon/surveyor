package cyclonedx

import (
	"github.com/mlmon/surveyor/source"
)

// See: https://cyclonedx.org/docs/1.6/json/

type SBOM struct {
	BomFormat    string     `json:"bomFormat"`
	SpecVersion  string     `json:"specVersion"`
	SerialNumber string     `json:"serialNumber"`
	Version      string     `json:"version"`
	Components   Components `json:"components"`
}

type Components []Component

type Component struct {
	// required fields
	Type ComponentType `json:"type"`
	Name string        `json:"name"`

	// optional fields
	Version string `json:"version"`
}

type ComponentType string

const Library ComponentType = "library"

func From(records *source.RecordSet) (*SBOM, error) {
	var sbom = SBOM{
		BomFormat:   "From",
		SpecVersion: "1.4",
		Version:     "1",
	}

	var components Components
	for _, record := range records.Records {
		if record.Source == source.PackageList {
			for _, entry := range record.Entries {
				components = append(components, Component{
					Name:    entry.Key,
					Version: entry.Value,
					Type:    Library,
				})
			}
		}
	}

	sbom.Components = components

	return &sbom, nil
}
