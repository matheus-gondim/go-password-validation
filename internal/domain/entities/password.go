package entities

type VerifyPassword struct {
	Password string `json:"password"`
	Rules    []Rule `json:"rules"`
}
