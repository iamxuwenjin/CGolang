package container

import (
	"fmt"
	"testing"
)

func TestVarArray(t *testing.T) {
	var a [3]int
	for _, i := range a {
		fmt.Println(i)
	}
	array := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	fmt.Println(array[1][1])
}
