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
		if err := command(array[0], array[1]); err != nil {
			os.Exit(1)
		}
	}
}

func command(subcommand string, data string) error {
	switch subcommand {
	case "insert":
		if err := Insert(data); err != nil {
			return err
		}
		fmt.Println("ok")
		return nil
	case "select":
		res := Select(data)
		fmt.Println(res.value, res.match)
		return nil
	}
	return errors.New("undifined command")
}
