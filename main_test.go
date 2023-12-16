package main

import (
	"fmt"
	"testing"
)

func Test_command(t *testing.T) {
	if err := command("put", "key1=foo"); err != nil {
		fmt.Printf("Unexpected error")
	}
	if err := command("get", "key1"); err != nil {
		fmt.Printf("Unexpected response")
	}
}
