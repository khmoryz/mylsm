package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/khmoryz/mylsm/memtable"
)

func main() {

	for {
		fmt.Print(">")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		array := strings.Split(scanner.Text(), " ")
		if err := command(array[0], array[1]); err != nil {
			os.Exit(1)
		}
	}
}

func command(subcommand string, data string) error {
	switch subcommand {
	case "put":
		if err := memtable.Put(data); err != nil {
			return err
		}
		fmt.Println("ok")
		return nil
	case "get":
		res := memtable.Get(data)
		fmt.Println(res.Value, res.Match)
		return nil
	case "del":
		if err := memtable.Delete(data); err != nil {
			return err
		}
		fmt.Println("ok")
		return nil
	}
	return errors.New("undifined command")
}
