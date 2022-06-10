package db

import (
	"database/sql"
	"fmt"

	"github.com/WeslleyRibeiro-1999/api-postgre/configs"
	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s database=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	conn, err := sql.Open("postgre", sc)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err

}
