package function

import (
	"fmt"
	"math"
)

func Typo(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func NameRetValue() (a, b int) {
	a = 1
	b = 2
	return
}

func Fire() {
	fmt.Println("this is fire")
}

func Visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}

func Accumulate(value int) func() int {
	return func() int {
		value ++
		return value
	}
}

func PlayerGen(name string) func() (string, int) {
	hp := 50
	return func() (string, int) {
		return name, hp
	}
}

func Args(args ...int)  {
	for _, v := range args{
		fmt.Println(v)
	}
}

func MultiType(args ...interface{})  {
	for _, v := range args {
		switch v.(type) {
		case int:
			fmt.Printf("%d is int\n", v)
		case string:
			fmt.Printf("%s is string\n", v)
		default:
			fmt.Print(v)
			fmt.Printf(" is unknown type \n")
		}
	}
}