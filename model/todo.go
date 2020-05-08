package model

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	UserID int    `json:"user_id"`
}
