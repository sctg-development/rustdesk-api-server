package dto

// User Registration
type UserRegReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	AuthKey  string `json:"auth_key"`
}

// The user changes the password
type UserSetPwdReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	AuthKey  string `json:"auth_key"`
}
