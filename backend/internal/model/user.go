package model

type AuthUser struct {
	Username     string
	Baid         uint
	PasswordHash string
}

type SimpleAuthUser struct {
	Username string
	Baid     uint
}
