package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3307)/golangtransacao")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	tx, _ := db.Begin()

	stmt, _ := tx.Prepare("INSERT INTO users(ID,NAME)VALUES(?,?)")
	stmt.Exec(1, "Danilo")
	stmt.Exec(2, "João")
	// _, err = stmt.Exec(1, "João")

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()

}
