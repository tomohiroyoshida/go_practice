package main

import (
  "database/sql"
	_ "github.com/lib/pq"
  "fmt"
)

var Db *sql.DB

func init() {
  var err error
  Db, err = sql.Open("postgres", "user=postgres dbname=postgres sslmode=disable")
  if err != nil {
    panic(err)
  }
}

// Get a single post
func retrieve(id int) (post Post, err error) {
  post := Post{}
  err := Db.QueryRow("SELECT name FROM users WHERE id=$1", id).Scan(&post.Id, &post.Content, &post.Author)
  return
}

func (post *Post) create() (err error) {
  statement := "insert into posts (content, author) values ($1, $2) returning id"
  stmt, err := Db.Prepare(statement)
  if err != nil {
    return
  }
  defer stmt.Close()
  err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
  return
}

func (post *Post) update() (err error) {
  _, err = Db.Exec("update posts set content = $1, author = $2 where id = $3", post.Content, post.Author, post.Id)
  return
}

func (post *Post) delete() (err error) {
  _, err = Db.Exec("delete from posts where id = $1", post.Id)
  return
}