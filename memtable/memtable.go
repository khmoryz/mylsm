package memtable

import (
	"fmt"
	"strings"
)

var Memt Memtable

type Memtable struct {
	Kvs []Kv
}

type Kv struct {
	Key   string
	Value string
}

type Result struct {
	Value string
	Match bool
}

func Insert(data string) error {
	d := strings.Split(data, "=")
	if len(d) != 2 {
		err := fmt.Errorf("unexpected insert value:%s", data)
		return err
	}
	kv := Kv{Key: d[0], Value: d[1]}
	Memt.Kvs = append(Memt.Kvs, kv)
	return nil
}

func Select(key string) Result {
	// Scan in reverse order (reason: because the latest value is appended at the end).
	for i := len(Memt.Kvs) - 1; i >= 0; i-- {
		if Memt.Kvs[i].Key == key {
			return Result{Value: Memt.Kvs[i].Value, Match: true}
		}

	}
	return Result{Value: "", Match: false}
}
