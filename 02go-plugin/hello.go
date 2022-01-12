package main

import "fmt"

// go build --buildmode=plugin -o hello.so hello.go
func Hello() {
	fmt.Println("Hello From plugin...")
}

func Test(a, b string) (c, d int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	c, d = 100, 200
	return c, d
}
