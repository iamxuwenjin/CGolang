package container

import (
	"fmt"
	"testing"
)

// 切片（slice）是对数组的一个连续片段的引用，所以切片是一个引用类型
func TestSlice(t *testing.T) {
	var numList []string
	fmt.Println(numList)

	a := make([]int, 2)
	b := make([]int, 2, 10)

	a = append(a, 1, 3, 4, 5)
	//当迭代切片时，关键字 range 会返回两个值，第一个值是当前迭代到的索引位置，第二个值是该位置对应元素值的一份副本
	for _, i := range a {
		fmt.Println(i)
	}
	fmt.Println(a, b)

}
