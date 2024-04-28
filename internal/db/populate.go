package db

import (
	"context"
	"database/sql"
	"time"
)

type Post struct {
	title   string
	content string
	author  string
}

var examplePosts = []Post{
	{"post 1", "this is the first post", "kytnacode"},
	{"post 2", "this is the second post", "kytnacode"},
	{"post 3", "this is the third post", "kytnacode"},
}

func PopulateDb(c context.Context, db *sql.DB) error {
	ctx, cancel := context.WithTimeout(c, time.Second*6)
	defer cancel()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		"CREATE TABLE IF NOT EXISTS posts (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT, content TEXT, author TEXT);",
	)
	if err != nil {
		return err
	}

	smft, err := tx.Prepare("INSERT INTO posts (title, content, author) VALUES (?, ?, ?);")
	if err != nil {
		return err
	}

	for _, post := range examplePosts {
		_, err = smft.Exec(post.title, post.content, post.author)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
