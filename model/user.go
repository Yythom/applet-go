package model

type UserRegisterParams struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

type UserInfoType struct {
	UserId           int    `json:"user_id"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	Gender           string `json:"gender"`
	Birthday         string `json:"birthday"`
	Address          string `json:"address"`
	LastLogin        string `json:"last_login"`
	AccountLocked    bool   `json:"account_locked"`
	RegistrationTime string `json:"registration_time"`
	UserType         string `json:"user_type"`
}
