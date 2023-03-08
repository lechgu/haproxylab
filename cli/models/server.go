package models

// Server ...
type Server struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Port    int    `json:"port"`
}

// ServersResponse ...
type ServersResponse struct {
	Servers []Server `json:"data"`
}

// AddServerRequest ...
type AddServerRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Port    int    `json:"port"`
}
