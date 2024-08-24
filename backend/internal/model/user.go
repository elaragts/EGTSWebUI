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

type UpdateAuthUser struct {
	Username        string `json:"username"`
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}
