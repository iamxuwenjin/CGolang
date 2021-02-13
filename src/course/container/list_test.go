package container

import (
	"container/list"
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	NodeList := list.New()
	NodeList.PushBack("canon")
	NodeList.PushFront("67")

	ele := NodeList.PushBack("1st")
	NodeList.InsertAfter("2nd", ele)

	for i := NodeList.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
