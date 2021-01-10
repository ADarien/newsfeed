package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

/*
create table if not exists "newsfeed" (
	"ID" integer not null primary key autoincrement,
	"content" text
);
*/

func main() {
	db, err := sql.Open("sqlite3", "./newsfeed.db")
}
