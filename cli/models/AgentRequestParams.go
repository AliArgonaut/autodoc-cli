package models

import (
	"encoding/json"
)

type AgentRequestParams struct {
	AppName     string          `json:"n"`
	Description string          `json:"d"`
	Filetree    json.RawMessage `json:"t"`
	Contents    json.RawMessage `json:"a"`
}
