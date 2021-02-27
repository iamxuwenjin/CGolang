package structure

import "fmt"

type class struct {
}

func (c *class) Do(v int) {
	fmt.Println("call method do:", v)
}

func FuncDo(v int) {
	fmt.Println("call method Do:", v)
}
