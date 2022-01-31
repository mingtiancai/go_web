package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Post struct {
	Id         int    `db:"id"`
	Content    string `db:"content"`
	AuthorName string `db:"author"`
}

func main() {
	db, err := sqlx.Connect("postgres", "user=gaoming dbname=t password=tiancai sslmode=disable")
	if err != nil {
		panic(err)
	}

	posts := []Post{}
	db.Select(&posts, "SELECT * FROM posts")
	for _, post := range posts {
		fmt.Println(post)
	}

}
