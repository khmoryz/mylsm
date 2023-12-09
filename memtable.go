package main

import (
	"fmt"
	"strings"
)

var Memt Memtable

type Memtable struct {
	kvs []Kv
}

type Kv struct {
	key   string
	value string
}

type Result struct {
	match bool
	value string
}

func Insert(data string) error {
	d := strings.Split(data, "=")
	if len(d) != 2 {
		err := fmt.Errorf("unexpected insert value:%s", data)
		return err
	}
	kv := Kv{key: d[0], value: d[1]}
	Memt.kvs = append(Memt.kvs, kv)
	return nil
}

func Select(key string) Result {
	for _, kv := range Memt.kvs {
		if kv.key == key {
			return Result{value: kv.value, match: true}
		}
	}
	return Result{value: "", match: false}
}
