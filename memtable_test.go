package main

import (
	"testing"
)

func TestInsertAndSelect(t *testing.T) {
	Insert("key=foo")
	if res := Select("key"); res.value != "foo" {
		t.Errorf("unexpected value:%s", res.value)
	}
}

func TestUpdate(t *testing.T) {
	Insert("key=foo")
	if res := Select("key"); res.value != "foo" {
		t.Errorf("unexpected value:%s", res.value)
	}
	Insert("key=bar")
	if res := Select("key"); res.value != "bar" {
		t.Errorf("unexpected value:%s", res.value)
	}
}
