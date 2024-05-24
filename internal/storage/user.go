package storage

type User struct {
	Id       int       `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Accounts []Account `json:"accounts"`
}
