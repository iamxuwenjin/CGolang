package variable

import (
	"fmt"
	"testing"
)

var (
	a int
	b string
)

// 变量命名
func TestVar(t *testing.T) {
	fmt.Println(a, b)
	c := 1
	d, e := "cat", "e"
	fmt.Println(c, d, e)
}

// 标准的变量命名方式
func TestSdtVar(t *testing.T) {
	// var name type = exps
	var hp int = 100
	// type 被认为是冗余信息，可简化为
	// var name = exps
	// 编译器会尝试根据右式类型推导左式类型
	var lp = 120
	fmt.Println(hp, lp)
}

func TestShortVar(t *testing.T) {
	// 未声明的情况下可用，且作用域为当前函数
	hp := 100
	fmt.Println(hp)
}

func getData() (int, int) {
	return 100, 200
}

func TestAnonymousVar(t *testing.T) {
	_, b := getData()
	a, _ := getData()
	fmt.Println(a, b)
}

// 声明全局变量
var B = "Bob"

func TestScope(t *testing.T) {
	// 声明局部变量
	var A = 3
	fmt.Println(A, B)
}
