package sstable

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

func TestRHeavy(t *testing.T) {
	cap := 10
	root := Insert(nil, "1")
	for i := 2; i < cap; i++ {
		fmt.Printf("=== Insert %d ===\n", i)
		root = Insert(root, strconv.Itoa(i))
		Print(root, 0)
	}
}

func TestLR(t *testing.T) {

	// default tree structure:
	//      6
	//     / \
	//    3   7
	//   / \
	//  2   4

	root := &AVLNode{Key: "6", Height: 4, Lhs: &AVLNode{Key: "3", Height: 3, Lhs: &AVLNode{Key: "2", Height: 1}, Rhs: &AVLNode{Key: "4", Height: 1}}, Rhs: &AVLNode{Key: "7", Height: 1}}
	fmt.Println("=== Default status ===")
	Print(root, 0)
	fmt.Println("=== Insert 5 ===")
	root = Insert(root, "5")

	// expected tree structure:
	//      4
	//     / \
	//    3   6
	//   /   / \
	//  2   5   7

	Print(root, 0)
}

func TestR(t *testing.T) {

	// tree structure:
	//      3
	//     /
	//    2

	root := &AVLNode{Key: "3", Height: 2, Lhs: &AVLNode{Key: "2", Height: 1}}
	fmt.Println("=== Default status ===")
	Print(root, 0)
	fmt.Println("=== Insert 1 ===")
	root = Insert(root, "1")

	// expected tree structure:
	//      2
	//     / \
	//    1   3
	Print(root, 0)
}

func TestRandomInsert(t *testing.T) {
	cap := 30
	root := Insert(nil, strconv.Itoa(rand.Int()))
	for i := 0; i < cap; i++ {
		root = Insert(root, strconv.Itoa(rand.Int()))
		Print(root, 0)
		fmt.Println("=======")
	}
	Print(root, 0)
}
