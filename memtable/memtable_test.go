package memtable

import (
	"testing"
)

func TestInsertAndSelect(t *testing.T) {
	Insert("key=foo")
	if res := Select("key"); res.Value != "foo" {
		t.Errorf("unexpected value:%s", res.Value)
	}
}

func TestUpdate(t *testing.T) {
	Insert("key=foo")
	if res := Select("key"); res.Value != "foo" {
		t.Errorf("unexpected value:%s", res.Value)
	}
	Insert("key=bar")
	if res := Select("key"); res.Value != "bar" {
		t.Errorf("unexpected value:%s", res.Value)
	}
}
