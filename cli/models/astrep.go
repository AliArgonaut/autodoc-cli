package models

type ASTFile struct {
	Name       string      `json:"f"`
	Signatures []Signature `json:"s"`
}

type Signature struct {
	Name    string   `json:"n"`
	Params  []string `json:"p"`
	Results []string `json:"r"`
}
