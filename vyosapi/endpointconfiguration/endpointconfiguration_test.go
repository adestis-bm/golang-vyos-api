package endpointconfiguration

import (
	"os"
	"reflect"
	"testing"
)

func TestSaveTo(t *testing.T) {
	filename := "/tmp/test.json"

	ec1 := &EndpointConfiguration{
		URL: "https://vyos/",
		Key: "MY-HTTPS-API-PLAINTEXT-KEY",
	}

	if err := ec1.SaveTo(filename); err != nil {
		t.Fatalf("SaveTo() failed with: %s", err)
	}
	defer os.Remove(filename)

	ec2, err := LoadFrom(filename)
	if err != nil {
		t.Fatalf("LoadFrom() failed with: %s", err)
	}

	if !reflect.DeepEqual(ec1, ec2) {
		t.Errorf("EndpointConfigurations do not match")
	}
}
