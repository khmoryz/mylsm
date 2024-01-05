package sstable

import (
	"fmt"
	"math"
)

type AVLNode struct {
	Key      string
	Height   int
	Lhs, Rhs *AVLNode
}

func setHeight(node *AVLNode) {
	if node.Lhs == nil && node.Rhs == nil {
		node.Height = 1
		return
	}
	if node.Lhs == nil && node.Rhs != nil {
		node.Height = node.Rhs.Height + 1
		return
	}
	if node.Rhs == nil && node.Lhs != nil {
		node.Height = node.Lhs.Height + 1
		return
	}

	node.Height = int(math.Max(float64(node.Lhs.Height), float64(node.Rhs.Height))) + 1
}

func Insert(node *AVLNode, key string) *AVLNode {
	if node == nil {
		// TODO: Height0のほうが良いかも
		return &AVLNode{Key: key, Height: 1}
	}

	if key < node.Key {
		node.Lhs = Insert(node.Lhs, key)
	} else if key > node.Key {
		node.Rhs = Insert(node.Rhs, key)
	}

	setHeight(node)
	r := rebalance(node)

	return r
}

func isBalanced(node *AVLNode) bool {
	if node.Lhs == nil {
		return node.Rhs.Height <= 1
	}
	if node.Rhs == nil {
		return node.Lhs.Height <= 1
	}
	return int(math.Abs(float64(node.Lhs.Height)-float64(node.Rhs.Height))) <= 1
}

func isLHeavy(node *AVLNode) bool {
	if node.Lhs == nil {
		return false
	}
	if node.Rhs == nil {
		return node.Lhs.Height == 1
	}
	return int(float64(node.Lhs.Height)-float64(node.Rhs.Height)) == 1
}

func bias(node *AVLNode) int {
	if node.Lhs == nil {
		return -node.Rhs.Height
	}
	if node.Rhs == nil {
		return node.Lhs.Height
	}
	return node.Lhs.Height - node.Rhs.Height
}

func rotateL(node *AVLNode) *AVLNode {
	if bias(node) == -2 && bias(node.Rhs) == -1 {
		newRoot := node.Rhs
		node.Rhs = node.Rhs.Lhs
		newRoot.Lhs = node

		setHeight(newRoot.Lhs)
		setHeight(newRoot)
		return newRoot
	}
	panic("tmp")
}

func rebalance(node *AVLNode) *AVLNode {
	if isBalanced(node) {
		return node
	}
	if isLHeavy(node) {
		//TODO
		fmt.Println("Unimplemented!")
		return nil
	} else {
		hoge := rotateL(node)
		return hoge
	}
}

func Print(root *AVLNode, indent int) {
	if root != nil {
		indent++
		Print(root.Rhs, indent)
		fmt.Printf("%*s", indent*2, "")
		fmt.Printf("|%s\n", root.Key)
		Print(root.Lhs, indent)
	}
}

var tree [][]string

func PrintVertical(root *AVLNode) {
	tree = [][]string{}
	genTree(root, 0, false)

	for _, t := range tree {
		fmt.Printf("%s\n", t)
	}
}

func genTree(root *AVLNode, depth int, isRhs bool) {
	if root == nil {
		return
	}

	for len(tree) <= depth {
		tree = append(tree, []string{})
	}

	tree[depth] = append(tree[depth], root.Key)
	if isRhs {
		tree[depth] = append(tree[depth], "|")
	}
	depth++
	genTree(root.Lhs, depth, false)
	genTree(root.Rhs, depth, true)
}
