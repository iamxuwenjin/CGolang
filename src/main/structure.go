package main

import (
	"fmt"
	"structure"
)

func main() {
	point := structure.Point{
		X: 4,
		Y: 5,
	}
	fmt.Println(point.Y)

	player := new(structure.Player)
	player.Hp = 66
	player.Mp = 88
	player.Name = "jojo"

	fmt.Println(player)

	res := structure.Mimi("jojo")
	fmt.Println(res)
}
