package interFace

import (
	"fmt"
	"testing"
)

func getType(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("the type of a is int")
	case string:
		fmt.Println("the type of a is string")
	case float64:
		fmt.Println("the type of a is float")
	default:
		fmt.Printf("unknow type")
	}
}

func TestAssert(t *testing.T) {
	var x interface{}
	x = 10
	value, ok := x.(int)
	fmt.Printf("value: %d, %t, type: %T\n", value, ok, value)

	getType("apple")
}
