package clients

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

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
	noBackends := []models.Backend{}
	client := http.Client{}
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/backends", haClient.baseURL)
	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		return noBackends, err
	}
	req.SetBasicAuth(haClient.username, haClient.password)
	resp, err := client.Do(req)
	if err != nil {
		return noBackends, nil
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return noBackends, nil
	}
	var backendResponse models.BackendsResponse
	err = json.Unmarshal(bodyBytes, &backendResponse)
	if err != nil {
		return noBackends, err
	}
	return backendResponse.Backends, nil
}

func (haClient *HaClient) ListServers(backend string) ([]models.Server, error) {
	noServers := []models.Server{}
	client := http.Client{}
	address := fmt.Sprintf("%s/v2/services/haproxy/configuration/servers", haClient.baseURL)
	req, err := http.NewRequest(http.MethodGet, address, http.NoBody)
	if err != nil {
		return noServers, err
	}
	q := url.Values{}
	q.Add("backend", backend)
	req.URL.RawQuery = q.Encode()

	req.SetBasicAuth(haClient.username, haClient.password)
	req.SetBasicAuth(haClient.username, haClient.password)
	resp, err := client.Do(req)
	if err != nil {
		return noServers, nil
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return noServers, nil
	}
	var serversResponse models.ServersResponse
	err = json.Unmarshal(bodyBytes, &serversResponse)
	if err != nil {
		return noServers, err
	}
	return serversResponse.Servers, nil
}
