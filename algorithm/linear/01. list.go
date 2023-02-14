package main

import (
	"container/list"
	"fmt"
)

// A list is a sequence of elements. Each element can be connected to another with a link
// in a forward or backward direction. The element can have other payload properties.
// This data structure is a basic type of container. Lists have a variable length and
// developer can remove or add elements more easily than an array. Data items within
// a list need not be contiguous in memory or on disk.

func main() {
	var intList list.List
	intList.PushBack(11)
	intList.PushBack(23)
	intList.PushBack(34)
	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}
}
