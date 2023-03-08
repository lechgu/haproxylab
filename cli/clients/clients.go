package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

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
	address := fmt.Sprintf("%s/v2/services/haproxy/configuration/backends", haClient.baseURL)
	req, err := http.NewRequest(http.MethodGet, address, http.NoBody)
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

// ListServers ...
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

// RemoveServer ...
func (haClient *HaClient) RemoveServer(server string, backend string) error {
	ver, err := haClient.GetConfigurationVersion()
	if err != nil {
		return err
	}
	address := fmt.Sprintf("%s/v2/services/haproxy/configuration/servers/%s", haClient.baseURL, server)
	req, err := http.NewRequest(http.MethodDelete, address, http.NoBody)
	if err != nil {
		return err
	}
	q := url.Values{}
	q.Add("backend", backend)
	q.Add("version", strconv.FormatInt(ver, 10))
	q.Add("force_reload", strconv.FormatBool(true))
	req.URL.RawQuery = q.Encode()
	req.SetBasicAuth(haClient.username, haClient.password)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bodyBytes))
	return err
}

// AddServer ...
func (haClient *HaClient) AddServer(name string, serverAddress string, port int, backend string) error {
	ver, err := haClient.GetConfigurationVersion()
	if err != nil {
		return err
	}
	address := fmt.Sprintf("%s//v2/services/haproxy/configuration/servers", haClient.baseURL)
	dto := models.AddServerRequest{
		Address: serverAddress,
		Name:    name,
		Port:    port,
	}
	payload, err := json.Marshal(dto)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(payload)
	req, err := http.NewRequest(http.MethodPost, address, reader)
	q := url.Values{}
	q.Add("backend", backend)
	q.Add("version", strconv.FormatInt(ver, 10))
	q.Add("force_reload", strconv.FormatBool(true))
	req.URL.RawQuery = q.Encode()
	req.SetBasicAuth(haClient.username, haClient.password)
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 201 {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(bodyBytes))
	}
	return err
}

// GetConfigurationVersion ...
func (haClient *HaClient) GetConfigurationVersion() (int64, error) {
	address := fmt.Sprintf("%s//v2/services/haproxy/configuration/version", haClient.baseURL)
	req, err := http.NewRequest(http.MethodGet, address, http.NoBody)
	if err != nil {
		return 0, err
	}
	req.SetBasicAuth(haClient.username, haClient.password)
	req.SetBasicAuth(haClient.username, haClient.password)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil
	}
	var ints = string(bodyBytes)
	return strconv.ParseInt(strings.TrimSpace(ints), 10, 64)
}
