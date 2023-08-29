package vyosapi

import (
	"encoding/json"
	"testing"

	"github.com/adestis-bm/golang-vyos-api/vyosapi/endpointconfiguration"
)

func TestVyOSAPIBasic(t *testing.T) {
	filename := "~/.vyos.d/vyos-tests.json"
	ep, err := endpointconfiguration.LoadFrom(filename)
	if err != nil {
		t.Fatalf("endpointconfiguration.LoadFrom() filed with: %s", err)
	}

	t.Logf("ep: %#v", ep)

	va := &VyOSAPI{
		Endpoint:  ep,
		UserAgent: "golang-vyos-api/test",
	}

	jsonResponseBytes, err := va.Retrieve("system", "login", "user", "bm")
	if err != nil {
		t.Fatalf("io.ReadAll() failed with: %s", err)
	}

	var jsonResponse any
	if err := json.Unmarshal(jsonResponseBytes, &jsonResponse); err != nil {
		t.Fatalf("json.Unmarshal() failed with: %s", err)
	}

	t.Logf("json: \n%#v", jsonResponse)
}
