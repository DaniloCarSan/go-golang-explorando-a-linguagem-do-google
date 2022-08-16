package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) (result sql.Result) {
	result, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}

	return
}

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3307)/")

	if err != nil {
		panic(err)
	}

	exec(db, "CREATE DATABASE IF NOT EXISTS cursogolang")
	exec(db, "USE cursogolang")
	exec(db, "DROP TABLE IF EXISTS users")
	exec(db, `CREATE TABLE users (
		ID INTEGER AUTO_INCREMENT,
		NAME VARCHAR(80),
		PRIMARY KEY (ID)
	)`)

	defer db.Close()
}
