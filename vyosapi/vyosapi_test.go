package vyosapi

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
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

	jsonRequestBytes := `{
		"op": "showConfig",
		"path": ["system","login","user"]
	}`

	formData := url.Values{
		"data": []string{jsonRequestBytes},
		"key":  []string{ep.Key},
	}

	httpposturl := fmt.Sprintf("%s/retrieve", ep.URL)

	request, err := http.NewRequest("POST", httpposturl, strings.NewReader(formData.Encode()))

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("User-Agent", "golang-vyos-api/test")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: ep.InsecureCertificate,
		},
	}

	client := &http.Client{
		Transport: tr,
	}
	response, err := client.Do(request)
	if err != nil {
		t.Fatalf("client.Do() failed with: %s", err)
	}
	defer response.Body.Close()

	t.Logf("status: %s", response.Status)

	jsonResponseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("io.ReadAll() failed with: %s", err)
	}

	var jsonResponse any
	if err := json.Unmarshal(jsonResponseBytes, &jsonResponse); err != nil {
		t.Fatalf("json.Unmarshal() failed with: %s", err)
	}

	t.Logf("json: \n%#v", jsonResponse)
}
