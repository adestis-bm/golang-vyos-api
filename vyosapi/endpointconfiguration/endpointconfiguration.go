package endpointconfiguration

import (
	"encoding/json"
	"os"

	"github.com/adestis-bm/golang-vyos-api/internal/pkg/utils"
)

// EndpointConfiguration hold necessarry values for VyOS API calls.
type EndpointConfiguration struct {
	// URL is the base for VyOS API calls.
	URL string `json:"url,omitempty"`
	// Key is the required access token.
	Key string `json:"key,omitempty"`
	//InsecureCertificate allows self-signed certificates.
	InsecureCertificate bool `json:"insecure_certificate,omitempty"`
}

// LoadFrom tries to read a JSON EndpointConfiguration specified by filename.
func LoadFrom(filename string) (*EndpointConfiguration, error) {
	fn, err := utils.Check(filename)
	if err != nil {
		return nil, err
	}

	bytes, err := os.ReadFile(fn)
	if err != nil {
		return nil, err
	}

	var ec EndpointConfiguration
	if err := json.Unmarshal(bytes, &ec); err != nil {
		return nil, err
	}

	return &ec, nil
}

// SaveTo tries to store an EndpointConfiguration as JSON specified by filename.
func (ec *EndpointConfiguration) SaveTo(filename string) error {
	bytes, err := json.MarshalIndent(ec, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, bytes, 0644)
}
