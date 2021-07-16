package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func main() {
	// (1) creating csv file
	csvFile, err := os.Create("posts.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	allPosts := []Post{
		{Id: 1, Content: "Hello World!", Author: "Sau Sheong"},
		{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"},
		{Id: 3, Content: "Hola Mundo!", Author: "Pedro"},
		{Id: 4, Content: "Hello World!", Author: "Sau Sheong"},
		{Id: 5, Content: "Bonjour Monde!", Author: "Pierre"},
		{Id: 6, Content: "Hola Mundo!", Author: "Pedro"},
	}

	writer := csv.NewWriter(csvFile) // creating writer
	// 書き込み作業
	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush() // データを確実に書き込むため

	// (2) reading a csv file
	file, err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var posts []Post
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post{Id: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}
	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Content)

}
