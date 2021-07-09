package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var n = flag.Bool("n", false, "通し番号を付与する")
	flag.Parse()
	var (
		files =  flag.Args()
		path, err = os.Executable()
	)
	// fmt.Println(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "失敗！", err)
	}
	// 実行ファイルのディレクトリ名
	path = filepath.Dir(path)
	fmt.Println(path)
	for x:=0; x<len(files); x++ {
		filepath, err := os.Open(filepath.Join(path, files[x]))
		if err != nil {
			fmt.Fprintln(os.Stderr, "読み込みに失敗", err)
		} else {
			scanner := bufio.NewScanner(filepath)
			for i := 1 ; scanner.Scan(); i++ {
				if *n {	// option `-n`
					fmt.Printf("%v", i)
				}
				fmt.Println(scanner.Text())
			}
		}
	}
}