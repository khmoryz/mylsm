package memtable

import (
	"testing"
)

func TestPutAndGet(t *testing.T) {
	Put("key=foo")
	if res := Get("key"); res.Value != "foo" {
		t.Errorf("unexpected value:%s", res.Value)
	}
}

func TestUpdate(t *testing.T) {
	Put("key=foo")
	if res := Get("key"); res.Value != "foo" {
		t.Errorf("unexpected value:%s", res.Value)
	}
	Put("key=bar")
	if res := Get("key"); res.Value != "bar" {
		t.Errorf("unexpected value:%s", res.Value)
	}
}
