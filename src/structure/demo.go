package structure

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

type Player struct {
	Name string
	Hp   int
	Mp   int
}

type Cat struct {
	Color string
	Name  string
}

func Mimi(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

type Wheel struct {
	Size int
}

type Engine struct {
	Power int
	Type  int
}
type Car struct {
	Wheel
	Engine
}

type Node struct {
	Data int
	Next *Node
}

type PreNode struct {
	Data int
	Pre  *PreNode
	Next *PreNode
}

func ShowNode(p *Node) {
	for p != nil {
		fmt.Println(*p)
		p = p.Next
	}
}
