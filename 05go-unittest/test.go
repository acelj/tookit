package src

import (
	"fmt"
	"strings"
)

func Add(a, b int) int {
	fmt.Println("a + b = ", a+b)
	return a + b
}

func Mul(a int, b int) int {
	fmt.Println("a * b = ", a*b)
	return a * b
}

func MyFunc(a, b string) string {
	re := strings.Join([]string{a, b, "LIU"}, "_")
	fmt.Println("re = ", re)
	return re
}
