package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	res, err := command(os.Args[1], os.Args[2])
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(res)
}

func command(subcommand string, data string) (string, error) {
	switch subcommand {
	case "insert":
		return "ok", nil
	case "select":
		return "hoge", nil
	}
	return "", errors.New("undifined command")
}
