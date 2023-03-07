package clients

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/lechgu/haproxylab/models"
)

// HaClient ...
type HaClient struct {
	baseURL  string
	username string
	password string
}

// New ...
func New(baseURL string, usename string, password string) *HaClient {
	return &HaClient{
		baseURL:  baseURL,
		username: usename,
		password: password,
	}
}

// ListBackends ...
func (haClient *HaClient) ListBackends() ([]models.Backend, error) {
	backends := []models.Backend{}
	client := http.Client{}
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/backends", haClient.baseURL)
	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		return backends, err
	}
	req.SetBasicAuth(haClient.username, haClient.password)
	resp, err := client.Do(req)
	if err != nil {
		return backends, nil
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return backends, nil
	}
	var backendResponse models.BackendResponse
	err = json.Unmarshal(bodyBytes, &backendResponse)
	if err != nil {
		return backends, err
	}
	return backendResponse.Backends, nil
}
