package vyosapi

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/adestis-bm/golang-vyos-api/vyosapi/endpointconfiguration"
)

type VyOSAPI struct {
	Endpoint  *endpointconfiguration.EndpointConfiguration
	UserAgent string
}

// https://docs.vyos.io/en/equuleus/automation/vyos-api.html#api-endpoints

func (va *VyOSAPI) post(path, data string) ([]byte, error) {
	formData := url.Values{
		"data": []string{data},
		"key":  []string{va.Endpoint.Key},
	}

	posturl := fmt.Sprintf("%s/%s", va.Endpoint.URL, path)
	request, err := http.NewRequest("POST", posturl, strings.NewReader(formData.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("User-Agent", va.UserAgent)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: va.Endpoint.InsecureCertificate,
		},
	}

	client := &http.Client{
		Transport: tr,
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return io.ReadAll(response.Body)
}

type RequestData struct {
	Operation string   `json:"op,omitempty"`
	Path      []string `json:"path,omitempty"`
}

func (va *VyOSAPI) Retrieve(path ...string) ([]byte, error) {
	data := RequestData{
		Operation: "showConfig",
		Path:      path,
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return va.post("retrieve", string(bytes))
}
