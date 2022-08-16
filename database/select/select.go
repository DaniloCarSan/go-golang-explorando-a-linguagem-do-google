package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	ID   int
	NAME string
}

func main() {

	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3307)/golangtransacao")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {

		var u user

		rows.Scan(&u.ID, &u.NAME)

		fmt.Println(u)

	}

}
