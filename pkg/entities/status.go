package entities

// Status - holds the information for the server's status
type Status struct {
	Status      string `json:"status"`
	Version     string `json:"version"`
	Environment string `json:"environment"`
}
