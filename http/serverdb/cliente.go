package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int    `json:"ID"`
	NAME string `json:"NAME"`
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/users/")
	id, _ := strconv.Atoi(sid)
	fmt.Println(r.URL.Path)
	switch {
	case r.Method == "GET" && id > 0:
		userById(w, r, id)
	case r.Method == "GET":
		userAll(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Sorry !")
	}
}

func userById(w http.ResponseWriter, r *http.Request, id int) {

	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3307)/golangtransacao")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var user User

	db.QueryRow("SELECT * FROM users WHERE ID = ?", id).Scan(&user.ID, &user.NAME)

	json, _ := json.Marshal(user)

	w.Header().Clone().Set("Content-Type", "application/json")

	fmt.Fprint(w, string(json))
}

func userAll(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3307)/golangtransacao")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var users []User

	rows, _ := db.Query("SELECT * FROM users")

	for rows.Next() {
		var user User

		rows.Scan(&user.ID, &user.NAME)

		users = append(users, user)
	}

	json, _ := json.Marshal(users)

	w.Header().Clone().Set("Content-Type", "application/json")

	fmt.Fprint(w, string(json))
}
