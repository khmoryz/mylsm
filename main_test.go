package main

import (
	"fmt"
	"testing"
)

func Test_command(t *testing.T) {
	if err := command("insert", "key1=foo"); err != nil {
		fmt.Printf("Unexpected error")
	}
	if err := command("select", "key1"); err != nil {

		fmt.Printf("Unexpected response")
	}
}
