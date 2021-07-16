package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data := []byte("Hello world!\n")

	// data1
	// ファイルへの書き込みとファイルからの読み込み。WriteFileとReadFileを利用
	err := ioutil.WriteFile("data1", data, 0644)
	if err != nil {
		panic(err)
	}
	read1, _ := ioutil.ReadFile("data1")
	fmt.Print(string(read1))

	// data2
	// File構造体を利用したファイルの読み書き
	file1, _ := os.Create("data2") // 作る
	defer file1.Close()

	bytes, _ := file1.Write(data) // 書き込む
	fmt.Printf("Wrote %d bytes to file\n", bytes)

	file2, _ := os.Open("data2") // 読み込む
	defer file2.Close()

	read2 := make([]byte, len(data))
	bytes, _ = file2.Read(read2)
	fmt.Printf("Read %d bytes from file\n", bytes)
	fmt.Println("data2: ", string(read2))
}
