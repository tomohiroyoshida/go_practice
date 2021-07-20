package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestDecode(t *testing.T) {
	post, err := decode("post.json")

	if err != nil {
		t.Error(err)
	}
	if post.Id != 1 {
		t.Error("Wrong id, was expecting 1 but got", post.Id)
	}
	if post.Content != "Hello World!" {
		t.Error("Wrong content, was expecting 'Hello World!' but got", post.Content)
	}
}

// target func
func decode(filename string) (post Post, err error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&post)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	return
}

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}
