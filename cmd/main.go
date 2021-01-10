package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "./newsfeed.db")

	stmt, _ := db.Prepare(`
		insert into newsfeed (content) values (?)
	`)

	stmt.Exec("Hola Youtube!")

	// stmt, _ := db.Prepare(`
	// 	create table if not exists "newsfeed" (
	// 		"ID" integer not null primary key autoincrement,
	// 		"content" text
	// 	);
	// `)

	// stmt.Exec()
}
