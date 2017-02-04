package models

type Account struct {
	Username string
	Password string
	AccountStatus
}

type AccountStatus struct {
	Banned  bool
	Active  bool
	Captcha bool
}
