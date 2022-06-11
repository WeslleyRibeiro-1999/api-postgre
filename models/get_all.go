package models

import "github.com/WeslleyRibeiro-1999/api-postgre/db"

func GetAll() (todos []Todo, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM todos`)
	if err != nil {
		return
	}

	for rows.Next() {
		var todo Todo

		err = rows.Scan(&todo.ID, &todo.Title, &todo.Descricao, &todo.Done)
		if err != nil {
			continue
		}

		todos = append(todos, todo)
	}

	return
}
