package models

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Email    int    `json:"email"`
	Password string `json:"password"`
	Admin    bool   `json:"admin"`
}
