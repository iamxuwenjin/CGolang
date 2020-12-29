package structure

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
