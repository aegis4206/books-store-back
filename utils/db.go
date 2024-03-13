package utils

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	serverInfo := "user=dbuser password=00000000 dbname=bookstore host=localhost port=5432 sslmode=disable"
	Db, err = sql.Open("postgres", serverInfo)
	if err != nil {
		panic(err.Error())
	}
}
