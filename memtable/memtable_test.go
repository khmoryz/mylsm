package memtable

import (
	"math/rand"
	"github.com/khmoryz/mylsm/sstable"
	"os"
	"testing"
)

func TestPutAndGet(t *testing.T) {
	// clean up
	os.RemoveAll(sstable.DirName)

	Put("key=foo")
	if res := Get("key"); res.Value != "foo" {
		t.Errorf("unexpected value:%s", res.Value)
	}
}

func TestUpdate(t *testing.T) {
	// clean up
	os.RemoveAll(sstable.DirName)

	Put("key=foo")
	if res := Get("key"); res.Value != "foo" {
		t.Errorf("unexpected value:%s", res.Value)
	}
	Put("key=bar")
	if res := Get("key"); res.Value != "bar" {
		t.Errorf("unexpected value:%s", res.Value)
	}
}

// note: Hight Cost.
func TestComplex(t *testing.T) {
	// clean up
	os.RemoveAll(sstable.DirName)

	const loop = 100

	for i := 0; i < loop; i++ {
		Put(RandomString(10) + "=" + RandomString(10))
	}
	Put("key1=foo")
	for i := 0; i < loop; i++ {
		Put(RandomString(10) + "=" + RandomString(10))
	}
	if res := Get("key1"); res.Value != "foo" {
		t.Errorf("unexpected value:%s, match:%t", res.Value, res.Match)
	}
}

func RandomString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
