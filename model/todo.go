package model

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Token string `json:"token"`
}
