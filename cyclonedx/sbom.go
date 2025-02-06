package cyclonedx

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
	Type string `json:"type"`
	Name string `json:"name"`

	// optional fields
	Version string `json:"version"`
}
