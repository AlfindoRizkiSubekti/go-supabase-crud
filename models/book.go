package models

type Book struct {
	ID         string     `json:"id"`
	Title      string     `json:"title"`
	Year       int        `json:"year"`
	Author     Author     `json:"author"`
	Categories []Category `json:"categories"`
}
