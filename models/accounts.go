package models

type NewAccount struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Account struct {
	Id       uint
	Username string
	Password string
}
