package resolver

type Token struct {
	TokenString string `json:"token"`
}

func (t *Token) Token() string {
	return t.TokenString
}
