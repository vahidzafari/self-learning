package main

import (
	"fmt"
	"sync"
)

// A binary search tree is a data structure that allows for the quick lookup, addition, and
// removal of elements. It stores the keys in a sorted order to enable a faster lookup. On
// average, space usage for a binary search tree is of the order O(n), whereas the insert, search,
// and delete operations are of the order O(log n). A binary search tree consists of nodes with
// properties or attributes:
// - A key integer
// - A value integer
// - The leftNode and rightNode instances of TreeNode

type TreeNode struct {
	key       int
	value     int
	leftNode  *TreeNode
	rightNode *TreeNode
}

type BinarySearchTree struct {
	rootNode *TreeNode
	lock     sync.RWMutex
}

func (tree *BinarySearchTree) InsertElement(key int, value int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	var treeNode *TreeNode
	treeNode = &TreeNode{key, value, nil, nil}

	if tree.rootNode == nil {
		tree.rootNode = treeNode
	} else {
		insertTreeNode(tree.rootNode, treeNode)
	}
}

// This method inserts the new TreeNode in the binary search tree.
func insertTreeNode(rootNode *TreeNode, newTreeNode *TreeNode) {
	if newTreeNode.key < rootNode.key {
		if rootNode.leftNode == nil {
			rootNode.leftNode = newTreeNode
		} else {
			insertTreeNode(rootNode.leftNode, newTreeNode)
		}
	} else {
		if rootNode.rightNode == nil {
			rootNode.rightNode = newTreeNode
		} else {
			insertTreeNode(rootNode.rightNode, newTreeNode)
		}
	}
}

func (tree *BinarySearchTree) InOrderTraverseTree(function func(int)) {
	tree.lock.RLock()
	defer tree.lock.RUnlock()

	inOrderTraverseTree(tree.rootNode, function)
}

// This method traverses the left, the root, and the right tree.
func inOrderTraverseTree(treeNode *TreeNode, function func(int)) {
	if treeNode != nil {
		inOrderTraverseTree(treeNode.leftNode, function)
		function(treeNode.value)
		inOrderTraverseTree(treeNode.rightNode, function)
	}
}

func (tree *BinarySearchTree) PreOrderTraverseTree(function func(int)) {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	preOrderTraverseTree(tree.rootNode, function)
}

func preOrderTraverseTree(treeNode *TreeNode, function func(int)) {
	if treeNode != nil {
		function(treeNode.value)
		preOrderTraverseTree(treeNode.leftNode, function)
		preOrderTraverseTree(treeNode.rightNode, function)
	}
}

func (tree *BinarySearchTree) PostOrderTraverseTree(function func(int)) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	postOrderTraverseTree(tree.rootNode, function)
}

func postOrderTraverseTree(treeNode *TreeNode, function func(int)) {
	if treeNode != nil {
		preOrderTraverseTree(treeNode.leftNode, function)
		preOrderTraverseTree(treeNode.rightNode, function)
		function(treeNode.value)
	}
}

// This method finds the node with the minimum value in the binary search tree.
func (tree *BinarySearchTree) MinNode() *int {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	var treeNode *TreeNode
	treeNode = tree.rootNode

	if treeNode == nil {
		return (*int)(nil)
	}

	for {
		if treeNode.leftNode == nil {
			return &treeNode.value
		}
		treeNode = treeNode.leftNode
	}
}

// This method finds the node with maximum property in the binary search tree.
func (tree *BinarySearchTree) MaxNode() *int {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	var treeNode *TreeNode
	treeNode = tree.rootNode

	if treeNode == nil {
		return (*int)(nil)
	}

	for {
		if treeNode.rightNode == nil {
			return &treeNode.value
		}
		treeNode = treeNode.rightNode
	}
}

func (tree *BinarySearchTree) SearchNode(key int) bool {
	tree.lock.RLock()
	defer tree.lock.RUnlock()

	return searchNode(tree.rootNode, key)
}

// The method returns true or false after checking whether treeNode with the same value as key exists
func searchNode(treeNode *TreeNode, key int) bool {
	if treeNode == nil {
		return false
	}
	if key < treeNode.key {
		return searchNode(treeNode.leftNode, key)
	}
	if key > treeNode.key {
		return searchNode(treeNode.rightNode, key)
	}
	return true
}

func (tree *BinarySearchTree) RemoveNode(key int) *TreeNode {
	tree.lock.RLock()
	defer tree.lock.RUnlock()

	return removeNode(tree.rootNode, key)
}

func removeNode(treeNode *TreeNode, key int) *TreeNode {
	if treeNode == nil {
		return nil
	}

	if key < treeNode.key {
		treeNode.leftNode = removeNode(treeNode.leftNode, key)
		return treeNode
	}

	if key < treeNode.key {
		treeNode.rightNode = removeNode(treeNode.rightNode, key)
		return treeNode
	}

	// key == node.key
	if treeNode.leftNode == nil && treeNode.rightNode == nil {
		treeNode = nil
		return nil
	}

	if treeNode.leftNode == nil {
		treeNode = treeNode.rightNode
		return treeNode
	}

	if treeNode.rightNode == nil {
		treeNode = treeNode.leftNode
		return treeNode
	}

	var leftmostrightNode *TreeNode
	leftmostrightNode = treeNode.rightNode

	for {
		// find smallest value on the right side
		if leftmostrightNode != nil && leftmostrightNode.leftNode != nil {
			leftmostrightNode = leftmostrightNode.leftNode
		} else {
			break
		}
	}

	treeNode.key, treeNode.value = leftmostrightNode.key,
		leftmostrightNode.value
	treeNode.rightNode = removeNode(treeNode.rightNode, treeNode.key)

	return treeNode

}

// This method turns the tree into a string format.
func (tree *BinarySearchTree) String() {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	fmt.Println("------------------------------------------------")
	stringify(tree.rootNode, 0)
	fmt.Println("------------------------------------------------")
}

func stringify(treeNode *TreeNode, level int) {
	if treeNode != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += " "
		}
		format += "---> "
		level++
		stringify(treeNode.leftNode, level)
		fmt.Printf(format+"%d\n", treeNode.key)
		stringify(treeNode.rightNode, level)
	}
}

func main() {
	var tree *BinarySearchTree = &BinarySearchTree{}
	tree.InsertElement(8, 8)
	tree.InsertElement(3, 3)
	tree.InsertElement(10, 10)
	tree.InsertElement(1, 1)
	tree.InsertElement(2, 2)
	tree.InsertElement(6, 6)
	tree.InsertElement(4, 4)
	tree.String()
}
