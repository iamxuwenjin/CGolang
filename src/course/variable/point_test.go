package variable

import (
	"fmt"
	"testing"
)

func TestPoint(t *testing.T) {
	// 通过&获取指针
	var cat = 1
	var str = "banana"
	fmt.Printf("%p %p\n", &cat, &str)

	// 通过指针获取值
	var house = "malibu point"
	ptr := &house
	fmt.Printf("ptr type: %T\n", ptr)
	fmt.Printf("ptr type: %p\n", ptr)
	value := *ptr
	fmt.Printf("value : %s\n", value)

}
