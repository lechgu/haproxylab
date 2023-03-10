package models

// Backend ...
type Backend struct {
	Name           string `json:"name"`
	Mode           string `json:"mode"`
	ConnectTimeout int    `json:"connect_timeout"`
	ServerTimeout  int    `json:"server_timeout"`
}

// BackendsResponse ...
type BackendsResponse struct {
	Backends []Backend `json:"data"`
}
