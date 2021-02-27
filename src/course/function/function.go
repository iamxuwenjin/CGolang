package function

import (
	"fmt"
	"strings"
	"time"
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

// 定义函数类型
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

// 内存缓存方法实现
const LIM = 45

var fibs [LIM]int

func FibonacciCache(n int) (res int) {
	if fibs[n] != 0 {
		res = fibs[n]
	}
	if n < 2 {
		return 1
	} else {
		res = FibonacciCache(n-1) + FibonacciCache(n-2)
	}
	fibs[n] = res
	return
}

// 计时器
func FuncDuration(num int, f func(num int) (res int)) {
	start := time.Now()
	f(num)
	elapsed := time.Since(start)
	fmt.Println("该函数执行完成耗时：", elapsed)
}
