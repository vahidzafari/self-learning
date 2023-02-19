package main

import (
	"encoding/json"
	"fmt"
)

// Adelson, Velski, and Landis pioneered the AVL tree data structure and hence it is named
// after them. It consists of height adjusting binary search trees. The balance factor is obtained
// by finding the difference between the heights of the left and right sub-trees. Balancing is
// done using rotation techniques. If the balance factor is greater than one, rotation shifts the
// nodes to the opposite of the left or right sub-trees. The search, addition, and deletion
// operations are processed in the order of O(log n).

// The KeyValue interface has the LessThan and EqualTo methods. The LessThan and
// EqualTo methods take KeyValue as a parameter and return a Boolean value after checking
// the less than or equal to condition.
type KeyValue interface {
	LessThan(KeyValue) bool
	EqualTo(KeyValue) bool
}

type integerKey int

func (k integerKey) LessThan(k1 KeyValue) bool {
	return k < k1.(integerKey)
}
func (k integerKey) EqualTo(k1 KeyValue) bool {
	return k == k1.(integerKey)
}

type TreeNode struct {
	KeyValue     KeyValue
	BalanceValue int
	LinkedNodes  [2]*TreeNode
}

// This method takes a node value and returns the opposite node's value.
func opposite(nodeValue int) int {
	return 1 - nodeValue
}

// This method rotates the node opposite to the specified sub-tree.
func singleRotation(rootNode *TreeNode, nodeValue int) *TreeNode {
	var saveNode *TreeNode
	saveNode = rootNode.LinkedNodes[opposite(nodeValue)]

	rootNode.LinkedNodes[opposite(nodeValue)] =
		saveNode.LinkedNodes[nodeValue]

	saveNode.LinkedNodes[nodeValue] = rootNode
	return saveNode
}

// This method rotates the node twice.
func doubleRotation(rootNode *TreeNode, nodeValue int) *TreeNode {
	var saveNode *TreeNode
	saveNode =
		rootNode.LinkedNodes[opposite(nodeValue)].LinkedNodes[nodeValue]
	rootNode.LinkedNodes[opposite(nodeValue)].LinkedNodes[nodeValue] =
		saveNode.LinkedNodes[opposite(nodeValue)]
	saveNode.LinkedNodes[opposite(nodeValue)] =
		rootNode.LinkedNodes[opposite(nodeValue)]
	rootNode.LinkedNodes[opposite(nodeValue)] = saveNode
	saveNode = rootNode.LinkedNodes[opposite(nodeValue)]
	rootNode.LinkedNodes[opposite(nodeValue)] = saveNode.LinkedNodes[nodeValue]
	saveNode.LinkedNodes[nodeValue] = rootNode
	return saveNode
}

// This method adjusts the balance of the tree.
func adjustBalance(rootNode *TreeNode, nodeValue int, balanceValue int) {
	var node *TreeNode
	node = rootNode.LinkedNodes[nodeValue]
	var oppNode *TreeNode
	oppNode = node.LinkedNodes[opposite(balanceValue)]
	switch oppNode.BalanceValue {
	case 0:
		rootNode.BalanceValue = 0
		node.BalanceValue = 0
	case balanceValue:
		rootNode.BalanceValue = -balanceValue
		node.BalanceValue = 0
	default:
		rootNode.BalanceValue = 0
		node.BalanceValue = balanceValue
	}
	oppNode.BalanceValue = 0
}

// This method changes the balance factor by a single or double rotation.
func BalanceTree(rootNode *TreeNode, nodeValue int) *TreeNode {
	var node *TreeNode
	node = rootNode.LinkedNodes[nodeValue]
	var balance int
	balance = 2*nodeValue - 1
	if node.BalanceValue == balance {
		rootNode.BalanceValue = 0
		node.BalanceValue = 0
		return singleRotation(rootNode, opposite(nodeValue))
	}
	adjustBalance(rootNode, nodeValue, balance)
	return doubleRotation(rootNode, opposite(nodeValue))
}

func insertRNode(rootNode *TreeNode, key KeyValue) (*TreeNode, bool) {
	if rootNode == nil {
		return &TreeNode{KeyValue: key}, false
	}
	var dir int
	dir = 0
	if rootNode.KeyValue.LessThan(key) {
		dir = 1
	}
	var done bool
	rootNode.LinkedNodes[dir], done = insertRNode(rootNode.LinkedNodes[dir],
		key)
	if done {
		return rootNode, true
	}
	rootNode.BalanceValue = rootNode.BalanceValue + (2*dir - 1)
	switch rootNode.BalanceValue {
	case 0:
		return rootNode, true
	case 1, -1:
		return rootNode, false
	}
	return BalanceTree(rootNode, dir), true
}

func InsertNode(treeNode **TreeNode, key KeyValue) {
	*treeNode, _ = insertRNode(*treeNode, key)
}

func RemoveNode(treeNode **TreeNode, key KeyValue) {
	*treeNode, _ = removeRNode(*treeNode, key)
}

func removeBalance(rootNode *TreeNode, nodeValue int) (*TreeNode, bool) {
	var node *TreeNode
	node = rootNode.LinkedNodes[opposite(nodeValue)]
	var balance int
	balance = 2*nodeValue - 1
	switch node.BalanceValue {
	case -balance:
		rootNode.BalanceValue = 0
		node.BalanceValue = 0
		return singleRotation(rootNode, nodeValue), false
	case balance:
		adjustBalance(rootNode, opposite(nodeValue), -balance)
		return doubleRotation(rootNode, nodeValue), false
	}
	rootNode.BalanceValue = -balance
	node.BalanceValue = balance
	return singleRotation(rootNode, nodeValue), true
}
func removeRNode(rootNode *TreeNode, key KeyValue) (*TreeNode, bool) {
	if rootNode == nil {
		return nil, false
	}
	if rootNode.KeyValue.EqualTo(key) {
		switch {
		case rootNode.LinkedNodes[0] == nil:
			return rootNode.LinkedNodes[1], false
		case rootNode.LinkedNodes[1] == nil:
			return rootNode.LinkedNodes[0], false
		}
		var heirNode *TreeNode
		heirNode = rootNode.LinkedNodes[0]
		for heirNode.LinkedNodes[1] != nil {
			heirNode = heirNode.LinkedNodes[1]
		}
		rootNode.KeyValue = heirNode.KeyValue
		key = heirNode.KeyValue
	}
	var dir int
	dir = 0
	if rootNode.KeyValue.LessThan(key) {
		dir = 1
	}
	var done bool
	rootNode.LinkedNodes[dir], done = removeRNode(rootNode.LinkedNodes[dir], key)
	if done {
		return rootNode, true
	}
	rootNode.BalanceValue = rootNode.BalanceValue + (1 - 2*dir)
	switch rootNode.BalanceValue {
	case 1, -1:
		return rootNode, true
	case 0:
		return rootNode, false
	}
	return removeBalance(rootNode, dir)
}
func main() {
	var treeNode *TreeNode
	fmt.Println("Tree is empty")
	var avlTree []byte
	avlTree, _ = json.MarshalIndent(treeNode, "", " ")
	fmt.Println(string(avlTree))
	fmt.Println("\n Add Tree")
	InsertNode(&treeNode, integerKey(5))
	InsertNode(&treeNode, integerKey(3))
	InsertNode(&treeNode, integerKey(8))
	InsertNode(&treeNode, integerKey(7))
	InsertNode(&treeNode, integerKey(6))
	InsertNode(&treeNode, integerKey(10))
	avlTree, _ = json.MarshalIndent(treeNode, "", " ")
	fmt.Println(string(avlTree))
	fmt.Println("\n Delete Tree")
	RemoveNode(&treeNode, integerKey(3))
	RemoveNode(&treeNode, integerKey(7))
	avlTree, _ = json.MarshalIndent(treeNode, "", " ")
	fmt.Println(string(avlTree))
}
