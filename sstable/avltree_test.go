package sstable

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

func TestIncrementalInsert(t *testing.T) {
	cap := 12
	root := Insert(nil, "1")
	for i := 2; i < cap; i++ {
		fmt.Printf("=== Insert %d ===\n", i)
		root = Insert(root, strconv.Itoa(i))
		Print(root, 0)
	}
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
