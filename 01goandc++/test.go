package main

import (
	"C"
	"fmt"
)

// 下面前面的空格不能有
//export Add
func Add(a, b int) int {
	return a + b
}

//export Foo
func Foo(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	c, d := 100, 200
	return c, d
}

// go build -o test.so -buildmode=c-shared test.go
func main() {
	fmt.Println("nihaoa")
	fmt.Println(Add(1, 4))
}
