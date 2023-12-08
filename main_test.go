package main

import (
	"fmt"
	"testing"
)

func Test_command(t *testing.T) {
	if res, err := command("insert", "key1=foo"); err != nil {
		fmt.Printf("Unexpected response: %s", res)
	}
	res, err := command("select", "key1")
	if err != nil || res != "foo" {
		fmt.Printf("Unexpected response: %s", res)
	}
}
