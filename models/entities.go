package models

type Todo struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Descricao string `json:"descricao"`
	Done      bool   `json:"done"`
}
