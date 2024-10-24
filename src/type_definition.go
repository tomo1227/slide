package main

import "fmt"

type Foo int

func main() {
	var a Foo
	var b int = 1
	a = b // コンパイルエラー
	fmt.Printf("%d", a)
}
