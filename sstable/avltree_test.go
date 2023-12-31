package sstable

// Node: Implemented for SStable indexing but not needed, will be used for Memtable.

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

func TestL(t *testing.T) {

	// tree structure:
	//      1
	//       \
	//        2

	root := &AVLNode{Key: "1", Height: 2, Lhs: &AVLNode{Key: "2", Height: 1}}
	fmt.Println("=== Default status ===")
	Print(root, 0)
	fmt.Println("=== Insert 3 ===")
	root = Insert(root, "3")

	// expected tree structure:
	//      2
	//     / \
	//    1   3
	Print(root, 0)

	if root.Rhs.Key != "3" {
		t.Errorf("Actual key is %s, expected 3", root.Rhs.Key)
	}
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
	if root.Rhs.Key != "3" {
		t.Errorf("Actual key is %s, expected 3", root.Rhs.Key)
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
	if root.Key != "4" {
		t.Errorf("Actual key is %s, expected 4", root.Key)
	}
}

func TestRL(t *testing.T) {

	// default tree structure:
	//
	//      3
	//     / \
	//    2   6
	//       / \
	//      5   7

	root := &AVLNode{Key: "3", Height: 3, Lhs: &AVLNode{Key: "2", Height: 1}, Rhs: &AVLNode{Key: "6", Height: 2, Lhs: &AVLNode{Key: "5", Height: 1}, Rhs: &AVLNode{Key: "7", Height: 1}}}
	fmt.Println("=== Default status ===")
	Print(root, 0)
	fmt.Println("=== Insert 4 ===")
	root = Insert(root, "4")

	// expected tree structure:
	//      5
	//     / \
	//    3   6
	//   / \   \
	//  2   4   7

	Print(root, 0)

	if root.Key != "5" {
		t.Errorf("Actual key is %s, expected 5", root.Key)
	}
}

func TestIncrementalInsert(t *testing.T) {
	cap := 30
	root := Insert(nil, "1")
	for i := 2; i < cap; i++ {
		root = Insert(root, strconv.Itoa(i))
	}
	Print(root, 0)
}

func TestRandomInsert(t *testing.T) {
	cap := 30
	root := Insert(nil, strconv.Itoa(rand.Int()))
	for i := 0; i < cap; i++ {
		root = Insert(root, strconv.Itoa(rand.Int()))
	}
	Print(root, 0)
}
