package structure

import "testing"

func TestFuncDo(t *testing.T) {
	var delegate func(int)
	c := new(class)

	delegate = c.Do
	delegate(100)

	delegate = FuncDo
	delegate(120)
}
