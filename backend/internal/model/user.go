package model

type AuthUser struct {
	Username     string `json:"username"`
	Baid         uint   `json:"baid"`
	PasswordHash string `json:"passwordHash"`
}

type SimpleAuthUser struct {
	Username string `json:"username"`
	Baid     uint   `json:"baid"`
}
