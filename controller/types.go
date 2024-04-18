package controller

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Group struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}
