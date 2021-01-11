package newsfeed

import "database/sql"

// Feed struct
type Feed struct {
	DB *sql.DB
}

// Get func
func (feed *Feed) Get() []Item {
	items := []Item{}
	rows, _ := feed.DB.Query(`
			select * from newsfeed
		`)
	var id int
	var content string
	for rows.Next() {
		rows.Scan(&id, &content)
		item := Item{
			ID:      id,
			Content: content,
		}
		items = append(items, item)
	}
	return items
}

// Add func
func (feed *Feed) Add(item Item) {
	stmt, _ := feed.DB.Prepare(`
		insert into newsfeed (content) values (?)
	`)

	stmt.Exec(item.Content)
}

// NewFeed func
func NewFeed(db *sql.DB) *Feed {
	stmt, _ := db.Prepare(`
		create table if not exists "newsfeed" (
			"ID" integer not null primary key autoincrement,
			"content" text
		);
	`)

	stmt.Exec()

	return &Feed{
		DB: db,
	}
}
