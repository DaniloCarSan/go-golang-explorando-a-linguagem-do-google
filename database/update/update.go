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

	stmt, _ := db.Prepare("UPDATE users SET name = ? WHERE ID = ?")

	stmt.Exec(1, "Danilo Carreio")

	stmt, _ = db.Prepare("DELETE FROM users WHERE ID = ?")
	stmt.Exec(1)

}
