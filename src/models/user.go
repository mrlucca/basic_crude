package models

type User struct {
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Age      int    `json:"age"`
	Password string `json:"password"`
}
