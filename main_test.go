package main

import (
	"fmt"
	"github.com/khmoryz/mylsm/sstable"
	"os"
	"testing"
)

func Test_command(t *testing.T) {
	// clean up
	os.RemoveAll(sstable.DirName)

	if err := command("put", "key1=foo"); err != nil {
		fmt.Printf("Unexpected error")
	}
	if err := command("get", "key1"); err != nil {
		fmt.Printf("Unexpected response")
	}
}
