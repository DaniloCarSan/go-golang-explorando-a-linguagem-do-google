package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3307)/cursogolang")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	stmt, _ := db.Prepare("INSERT INTO users(NAME)VALUES(?)")
	stmt.Exec("Danilo")
	stmt.Exec("Maria")

	res, _ := stmt.Exec("Pedro")

	id, _ := res.LastInsertId()

	fmt.Println("ID PEDRO:", id)

	rows, _ := res.RowsAffected()
	fmt.Println(rows)

}
