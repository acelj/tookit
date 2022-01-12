package main

import (
	"fmt"
	"os"
	"plugin"
)

// go run main.go
func main() {
	p, err := plugin.Open("./hello.so")
	if err != nil {
		fmt.Println("error open plugin: ", err)
		os.Exit(-1)
	}
	s, err := p.Lookup("Hello")
	if err != nil {
		fmt.Println("error lookup Helllo: ", err)
		os.Exit(-1)
	}
	if myhello, ok := s.(func()); ok {
		myhello()
	}

	s1, err := p.Lookup("Test")
	if err != nil {
		fmt.Println("error lookup Test: ", err)
		os.Exit(-1)
	}
	if myfunc, ok := s1.(func(string, string) (int, int)); ok {
		a, b := myfunc("111", "222")
		fmt.Printf("a = %d, b = %d\n", a, b)
	}
}
