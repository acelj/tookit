package src

import (
	"fmt"
	"testing"
)

// 文件名必须是_test.go 结尾
func TestAdd(t *testing.T) {
	fmt.Println("this is Testing")
	re := Add(1, 2)
	fmt.Println("re = ", re)
	if ans := Mul(1, 3); ans != 3 {
		t.Errorf("error, ans = %d", ans)
	}
}

func TestMyFunc(t *testing.T) {
	fmt.Println("This is MyFunc testing")
	re := MyFunc("de", "HUA")
	fmt.Println("re = ", re)
}
