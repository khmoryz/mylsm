package sstable

import (
	"math/rand"
	"os"
	"strconv"
	"testing"
)

func TestFlushAndSearch(t *testing.T) {
	// clean up
	os.RemoveAll(DirName)

	// set dummy data
	Memt.Kvs = []Kv{
		{Key: "a", Value: "foo"},
		{Key: "c", Value: "bar"},
		{Key: "b", Value: "buzz"},
		{Key: "a", Value: "foo_updated"},
	}

	// execute
	Flush(&Memt)

	// check
	v, exist := Search("a")
	if v != "foo_updated" && exist != true {
		t.Errorf("unexpected value: %s, exist: %t", v, exist)
	}
}

func TestDoubleFlushAndSearch(t *testing.T) {
	// clean up
	os.RemoveAll(DirName)

	// set first dummy data
	Memt.Kvs = []Kv{
		{Key: "a", Value: "foo"},
		{Key: "c", Value: "bar"},
		{Key: "b", Value: "buzz"},
		{Key: "a", Value: "foo_updated"},
	}

	// execute
	Flush(&Memt)

	// set first dummy data
	Memt.Kvs = []Kv{
		{Key: "x", Value: "foo"},
		{Key: "y", Value: "bar"},
		{Key: "z", Value: "buzz"},
		{Key: "b", Value: "buzz_updated_in_the_second"},
	}

	// execute
	Flush(&Memt)

	// check
	v, exist := Search("a")
	if v != "foo_updated" && exist != true {
		t.Errorf("unexpected value: %s, exist: %t", v, exist)
	}

	v, exist = Search("x")
	if v != "foo" && exist != true {
		t.Errorf("unexpected value: %s, exist: %t", v, exist)
	}

	v, exist = Search("b")
	if v != "buzz_updated_in_the_second" && exist != true {
		t.Errorf("unexpected value: %s, exist: %t", v, exist)
	}
}

func TestRandomWrite(t *testing.T) {
	// clean up
	os.RemoveAll(DirName)

	// set data
	cap := 30
	for i := 0; i < cap; i++ {
		Memt.Kvs = append(Memt.Kvs, Kv{Key: strconv.Itoa(rand.Int()), Value: strconv.Itoa(rand.Int())})
	}
	Memt.Kvs = append(Memt.Kvs, Kv{Key: "a", Value: "foo"})
	for i := 0; i < cap; i++ {
		Memt.Kvs = append(Memt.Kvs, Kv{Key: strconv.Itoa(rand.Int()), Value: strconv.Itoa(rand.Int())})
	}

	// execute
	Flush(&Memt)

	// check: existing data can be retrieved.
	v, exist := Search("a")
	if v != "foo" && exist != true {
		t.Errorf("unexpected value: %s, exist: %t", v, exist)
	}

	// check: non-existent data can't be retrieved.
	v, exist = Search("null!!")
	if v != "" && exist != false {
		t.Errorf("unexpected value: %s, exist: %t", v, exist)
	}
}
