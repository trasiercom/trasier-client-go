package api

type Endpoint struct {
	Name string `json:"name"`
	IpAddress string `json:"ipAddress"`
	Port string `json:"port"`
	Hostname string `json:"hostname"`
}
