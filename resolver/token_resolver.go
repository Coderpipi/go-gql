package resolver

type Token struct {
	TokenString string        `json:"token"`
	UserInfo    *UserResolver `json:"user"`
}

func (t *Token) Token() string {
	return t.TokenString
}

func (t *Token) User() *UserResolver {
	return t.UserInfo
}
