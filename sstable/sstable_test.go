package sstable

import (
	"fmt"
	"mylsm/memtable"
	"os"
	"testing"
)

func TestFlush(t *testing.T) {
	// clean up
	os.RemoveAll(dirName)

	// set dummy data
	memtable.Memt.Kvs = []memtable.Kv{
		{Key: "a", Value: "foo"},
		{Key: "c", Value: "hoge"},
		{Key: "b", Value: "bar"},
		{Key: "a", Value: "foo_updated"},
	}

	// execute
	Flush()

	// check
	v, exist := Read("a")
	fmt.Println(v, exist)
	if v != "foo_updated" && exist != true {
		t.Errorf("unexpected value: %s, exist: %t", v, exist)
	}
}
