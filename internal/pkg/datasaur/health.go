package datasaur

type HealthStatus struct {
	Status string `json:"status"`
}

type HealthResponse struct {
	Version string                  `json:"version"`
	Status  string                  `json:"status"`
	Details map[string]HealthStatus `json:"details"`
}
