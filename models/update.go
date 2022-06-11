package models

import "github.com/WeslleyRibeiro-1999/api-postgre/db"

func Update(id int64, todo Todo) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	res, err := conn.Exec(`UPDATE todos SET title=$2, description=$3, done=%4 WHERE id=$1`, id, todo.Title, todo.Descricao, todo.Done)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
