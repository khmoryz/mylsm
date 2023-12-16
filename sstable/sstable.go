package sstable

import (
	"encoding/binary"
	"io"
	"mylsm/memtable"
	"os"
	"sort"
)

const dirName = "/tmp/mylsm/"
const fileName = "sstable.mylsm"

type SSTable struct {
	kvs []kv
}

type kv struct {
	key   string
	value string
}

func Flush() error {
	//TODO:memtable exclusive control

	// uniq
	uniq := make(map[string]string, 0)
	for _, k := range memtable.Memt.Kvs {
		uniq[k.Key] = k.Value
	}

	// sort
	var sst SSTable
	for i, v := range uniq {
		sst.kvs = append(sst.kvs, kv{key: i, value: v})
	}
	sort.Slice(sst.kvs, func(i, j int) bool { return sst.kvs[i].key < sst.kvs[j].key })

	// write

	if err := os.MkdirAll(dirName, 0755); err != nil {
		panic(err)
	}

	f, err := os.OpenFile(dirName+fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for _, v := range sst.kvs {
		binary.Write(f, binary.LittleEndian, int32(len(v.key)))
		f.Write([]byte(v.key))
		binary.Write(f, binary.LittleEndian, int32(len(v.value)))
		f.Write([]byte(v.value))
	}

	return nil
}

func Read(searchKey string) (string, bool) {
	f, err := os.Open(dirName + fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for {
		// Read key
		var klen int32
		if err := binary.Read(f, binary.LittleEndian, &klen); err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		k := make([]byte, klen)
		if _, err := f.Read(k); err != nil {
			panic(err)
		}

		// Read value
		// TODO: Skipping read value if the key matches.
		var vlen int32
		if err := binary.Read(f, binary.LittleEndian, &vlen); err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		v := make([]byte, vlen)
		if _, err := f.Read(v); err != nil {
			panic(err)
		}

		// compare
		if string(k) == searchKey {
			return string(v), true
		}
	}

	return "", false
}
