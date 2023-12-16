package memtable

import (
	"fmt"
	st "mylsm/sstable"
	"strings"
)

var memtMax = 3

type Result struct {
	Value string
	Match bool
}

func Put(data string) error {
	d := strings.Split(data, "=")
	if len(d) != 2 {
		err := fmt.Errorf("unexpected insert value:%s", data)
		return err
	}
	kv := st.Kv{Key: d[0], Value: d[1]}
	st.Memt.Kvs = append(st.Memt.Kvs, kv)

	// Flush to sstable.
	if len(st.Memt.Kvs) >= memtMax {
		if err := st.Flush(); err != nil {
			return err
		}
	}

	return nil
}

func Get(key string) Result {
	// Scan in reverse order (reason: because the latest value is appended at the end).
	for i := len(st.Memt.Kvs) - 1; i >= 0; i-- {
		if st.Memt.Kvs[i].Key == key {
			return Result{Value: st.Memt.Kvs[i].Value, Match: true}
		}
	}

	// Search sstable.
	if v, m := st.Search(key); m {
		return Result{Value: v, Match: m}
	}

	return Result{Value: "", Match: false}
}
