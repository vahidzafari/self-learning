package main

import (
	"container/heap"
	"fmt"
)

// A heap is a data structure that is based on the heap property. The heap data structure is
// used in selection, graph, and k-way merge algorithms. Operations such as finding,
// merging, insertion, key changes, and deleting are performed on heaps. Heaps are part of
// the container/heap package in Go. According to the heap order (maximum heap)
// property, the value stored at each node is greater than or equal to its children.
// If the order is descending, it is referred to as a maximum heap; otherwise, it's a minimum
// heap.

type IntegerHeap []int

// IntegerHeap method - gets the length of integerHeap
func (iheap IntegerHeap) Len() int { return len(iheap) }

// IntegerHeap method - checks if element of i index is less than j index
func (iheap IntegerHeap) Less(i, j int) bool { return iheap[i] < iheap[j] }

// IntegerHeap method - swaps the element of i to j index
func (iheap IntegerHeap) Swap(i, j int) {
	iheap[i], iheap[j] = iheap[j],
		iheap[i]
}

// IntegerHeap method -pushes the item
func (iheap *IntegerHeap) Push(heapintf interface{}) {
	*iheap = append(*iheap, heapintf.(int))
}

// IntegerHeap method -pops the item from the heap
func (iheap *IntegerHeap) Pop() interface{} {
	var n int
	var x1 int
	var previous IntegerHeap = *iheap
	n = len(previous)
	x1 = previous[n-1]
	*iheap = previous[0 : n-1]
	return x1
}

func main() {
	var intHeap *IntegerHeap = &IntegerHeap{1, 4, 5}
	heap.Init(intHeap)
	heap.Push(intHeap, 2)
	fmt.Printf("minimum: %d\n", (*intHeap)[0])
	for intHeap.Len() > 0 {
		fmt.Printf("%d \n", heap.Pop(intHeap))
	}
}
