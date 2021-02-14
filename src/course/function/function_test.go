package function

import (
	"fmt"
	"testing"
)

func TestVisit(t *testing.T) {
	list := []int{1, 3, 5}
	Visit(list, func(num int) {
		fmt.Println(num)
	})
}

func TestFuncCaller(t *testing.T) {
	// 声明接口变量
	var invoker Invoker
	// 实例结构体
	s := new(Struct)
	// 实例化的结构体赋值到接口
	invoker = s
	invoker.Call("hello")

	// 将匿名函数转换为FuncCaller类型，在赋值给接口
	invoker = FuncCaller(func(v interface{}) {
		fmt.Println("from function", v)
	})
	invoker.Call("hello")
}

func TestJoinStrings(t *testing.T) {
	fmt.Println(JoinStrings("pig", "and", "rat"))
}

func TestFibonacci(t *testing.T) {
	fmt.Println(Fibonacci(10))

}
