package cyclonedx

import (
	"github.com/mlmon/surveyor/source"
)

// Nvidia nSpect uses 1.4 of the spec:
// https://cyclonedx.org/docs/1.4/json/

type SBOM struct {
	BomFormat    string     `json:"bomFormat"`
	SpecVersion  string     `json:"specVersion"`
	SerialNumber string     `json:"serialNumber"`
	Version      string     `json:"version"`
	Components   Components `json:"components"`
	Metadata     Metadata   `json:"metadata"`
}

type Metadata struct {
	Properties Properties `json:"properties"`
}

type Properties []Property

type Property struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Components []Component

type Component struct {
	// required fields
	Type ComponentType `json:"type"`
	Name string        `json:"name"`

	// optional fields
	Version string `json:"version"`
}

// ComponentType defined in 1.4 spec.
type ComponentType string

const Application ComponentType = "application"
const Container ComponentType = "container"
const Device ComponentType = "device"
const File ComponentType = "file"
const Firmware ComponentType = "firmware"
const Framework ComponentType = "framework"
const Library ComponentType = "library"
const OperatingSystem ComponentType = "operating-system"

var map_nvidia_smi = map[string]string{
	"vbios_version":  "vbios",
	"driver_version": "driver",
}

func From(records *source.RecordSet) (*SBOM, error) {
	uuid, err := Uuid()
	if err != nil {
		return nil, err
	}

	var sbom = SBOM{
		BomFormat:    "CycloneDX",
		SpecVersion:  "1.4",
		Version:      "1",
		SerialNumber: "urn:uuid:" + uuid,
	}

	var components Components
	for _, record := range records.Records {
		var componentType = Library

		if record.Source == source.Procfs {
			componentType = File
		} else if record.Source == "uname" {
			componentType = OperatingSystem
		}

		for _, entry := range record.Entries {
			name := entry.Key
			version := entry.Value

			if record.Source == "nvidia-smi" {
				if name == "name" {
					name = "gpu"
					componentType = Device
				} else {
					n, ok := map_nvidia_smi[name]
					if ok {
						name = n
					}
					componentType = Firmware
				}
			}

			components = append(components, Component{
				Name:    name,
				Version: version,
				Type:    componentType,
			})
		}
	}

	sbom.Components = components

	return &sbom, nil
}
