package main

import "fmt"

// A Set is a linear data structure that has a collection of values that are not repeated. A set can
// store unique values without any particular order. In the real world, sets can be used to
// collect all tags for blog posts and conversation participants in a chat. The data can be of
// Boolean, integer, float, characters, and other types. Static sets allow only query operations,
// which means operations related to querying the elements. Dynamic and mutable sets allow
// for the insertion and deletion of elements. Algebraic operations such as union, intersection,
// difference, and subset can be defined on the sets.

type Set struct {
	integerMap map[int]bool
}

func (set *Set) New() {
	set.integerMap = make(map[int]bool)
}

func (set *Set) AddElement(element int) {
	if !set.ContainsElement(element) {
		set.integerMap[element] = true
	}
}

func (set *Set) DeleteElement(element int) {
	delete(set.integerMap, element)
}

func (set *Set) ContainsElement(element int) bool {
	var exists bool
	_, exists = set.integerMap[element]

	return exists
}

func (set *Set) Intersect(anotherSet *Set) *Set {
	var intersectSet = &Set{}
	intersectSet.New()
	var value int

	for value, _ = range set.integerMap {
		if anotherSet.ContainsElement(value) {
			intersectSet.AddElement(value)
		}
	}
	return intersectSet
}

func (set *Set) Union(anotherSet *Set) *Set {
	var unionSet = &Set{}
	unionSet.New()
	var value int

	for value, _ = range set.integerMap {
		unionSet.AddElement(value)
	}

	for value, _ = range anotherSet.integerMap {
		unionSet.AddElement(value)
	}

	return unionSet
}

func main() {

	var set *Set
	set = &Set{}

	set.New()

	set.AddElement(1)
	set.AddElement(2)

	fmt.Println("initial set", set)
	fmt.Println(set.ContainsElement(1))

	var anotherSet *Set
	anotherSet = &Set{}

	anotherSet.New()
	anotherSet.AddElement(2)
	anotherSet.AddElement(4)
	anotherSet.AddElement(5)

	fmt.Println("another set", set)

	fmt.Println("intersection of sets ", set.Intersect(anotherSet))

	fmt.Println("union of sets ", set.Union(anotherSet))

}
