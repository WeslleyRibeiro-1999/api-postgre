package models

import (
	"github.com/WeslleyRibeiro-1999/api-postgre/db"
)

func Insert(todo Todo) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO todo (title, description, done) VALUES ($1, $2, $3) RETURN id`

	err = conn.QueryRow(sql, todo.Title, todo.Descricao, todo.Done).Scan(&id)

	return
}
