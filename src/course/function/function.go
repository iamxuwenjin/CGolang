package function

import (
	"fmt"
	"strings"
)

// 匿名函数作为回调函数
func Visit(list []int, f func(num int)) {
	for _, v := range list {
		f(v)
	}
}

// 函数实现接口
type Invoker interface {
	Call(interface{})
}

// 结构体类型
type Struct struct{}

func (s *Struct) Call(p interface{}) {
	fmt.Println("from struct", p)
}

// 函数定义为类型
type FuncCaller func(interface{})

// 实现Invoker的Call
func (f FuncCaller) Call(p interface{}) {
	// 调用函数主体
	f(p)
}

//遍历可变参数列表
func JoinStrings(list ...string) string {
	// 定义一个字节缓存，快速链接字符串
	var b strings.Builder
	for _, str := range list {
		b.WriteString(str + " ")
	}
	return b.String()
}

// 递归函数
func Fibonacci(n int) (res int) {
	if n < 2 {
		return 1
	} else {
		res = Fibonacci(n-1) + Fibonacci(n-2)
	}
	return
}
