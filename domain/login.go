package domain

type UserRegisterParams struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}
type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
