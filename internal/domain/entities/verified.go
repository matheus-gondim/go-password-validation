package entities

type Verified struct {
	Verify  bool     `json:"verify"`
	NoMatch []string `json:"noMatch"`
}
