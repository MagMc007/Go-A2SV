package models

type Task struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Duedate string `json:"duedate"`
	Status string `json:"status"`
}