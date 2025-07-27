package simulator

type SimulatorConfig struct {
	Endpoints []Endpoint `json:"endpoints"`
}

type Endpoint struct {
	Path    string            `json:"path"`
	Body    string            `json:"body"`
	Headers map[string]string `json:"headers"`
	Status  int               `json:"status"`
}

type ConfiguredResponse struct {
	Endpoint *Endpoint
	// params & etc
}
