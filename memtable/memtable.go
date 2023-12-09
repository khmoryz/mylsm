package memtable

import (
	"fmt"
	"strings"
)

var Memt Memtable

type Memtable struct {
	kvs []kv
}

type kv struct {
	key   string
	value string
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
	kv := kv{key: d[0], value: d[1]}
	Memt.kvs = append(Memt.kvs, kv)
	return nil
}

func Select(key string) Result {
	// Scan in reverse order (reason: because the latest value is appended at the end).
	for i := len(Memt.kvs) - 1; i >= 0; i-- {
		if Memt.kvs[i].key == key {
			return Result{Value: Memt.kvs[i].value, Match: true}
		}

	}
	return Result{Value: "", Match: false}
}
