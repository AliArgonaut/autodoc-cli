package models

import (
	"encoding/json"
)

type AgentRequestParams struct {
	AppName     string          `json:"n"`
	Description string          `json:"d"`
	Filetree    json.RawMessage `json:"t"`
	AST         json.RawMessage `json:"a"`
}
