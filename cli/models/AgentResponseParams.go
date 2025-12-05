package models

type AgentResponseParams struct {
	Success       bool   `json:"success"`
	Documentation string `json:"documentation"`
	Error         string `json:"error,omitempty"`
}
