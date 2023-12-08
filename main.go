package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {

	for {
		fmt.Print(">")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		array := strings.Split(scanner.Text(), ":")
		res, err := command(array[0], array[1])
		if err != nil {
			os.Exit(1)
		}
		fmt.Println(res)
	}
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
