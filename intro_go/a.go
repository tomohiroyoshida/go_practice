package main

import "fmt"

func main() {
	var s1 = make([]int, 3, 10)
	fmt.Println(s1)
}

	// interface と埋め込み
	type Hoge interface {M(); N()}
	type fuga struct {Hoge}
	func (f fuga) M() {
		fmt.Println("HI")
		f.Hoge.M()
	}
	func HiHoge (h Hoge) Hoge {
		return fuga{h}
	}