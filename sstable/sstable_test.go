package sstable

import (
	"os"
	"testing"
)

func TestFlushAndSearch(t *testing.T) {
	// clean up
	os.RemoveAll(dirName)

	// set dummy data
	Memt.Kvs = []Kv{
		{Key: "a", Value: "foo"},
		{Key: "c", Value: "bar"},
		{Key: "b", Value: "buzz"},
		{Key: "a", Value: "foo_updated"},
	}

	// execute
	Flush()

	// check
	v, exist := Search("a")
	if v != "foo_updated" && exist != true {
		t.Errorf("unexpected value: %s, exist: %t", v, exist)
	}
}

func TestDoubleFlushAndSearch(t *testing.T) {
	// clean up
	os.RemoveAll(dirName)

	// set first dummy data
	Memt.Kvs = []Kv{
		{Key: "a", Value: "foo"},
		{Key: "c", Value: "bar"},
		{Key: "b", Value: "buzz"},
		{Key: "a", Value: "foo_updated"},
	}

	// execute
	Flush()

	// set first dummy data
	Memt.Kvs = []Kv{
		{Key: "x", Value: "foo"},
		{Key: "y", Value: "bar"},
		{Key: "z", Value: "buzz"},
		{Key: "b", Value: "buzz_updated_in_the_second"},
	}

	// execute
	Flush()

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
