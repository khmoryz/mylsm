package sstable

import (
	"encoding/binary"
	"io"
	"io/fs"
	"os"
	"sort"
	"strconv"
	"time"
)

const dirName = "/tmp/mylsm/"
const suffix = "_sstable.mylsm"
const maxSstableSize = 1000

var Memt Table

type Table struct {
	Kvs []Kv
}

type Kv struct {
	Key   string
	Value string
}

func genFileName() string {
	now := time.Now().UnixNano()
	return strconv.FormatInt(now, 10) + suffix
}

func Flush() error {
	// uniq
	uniq := make(map[string]string, 0)
	for _, k := range Memt.Kvs {
		uniq[k.Key] = k.Value
	}

	// sort
	var sst Table
	for i, v := range uniq {
		sst.Kvs = append(sst.Kvs, Kv{Key: i, Value: v})
	}
	sort.Slice(sst.Kvs, func(i, j int) bool { return sst.Kvs[i].Key < sst.Kvs[j].Key })

	// write
	if err := os.MkdirAll(dirName, 0755); err != nil {
		panic(err)
	}

	f, err := os.OpenFile(dirName+genFileName(), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for _, v := range sst.Kvs {
		binary.Write(f, binary.LittleEndian, int32(len(v.Key)))
		f.Write([]byte(v.Key))
		binary.Write(f, binary.LittleEndian, int32(len(v.Value)))
		f.Write([]byte(v.Value))
	}

	// Initialization
	Memt = Table{}
	
	return nil
}

func getSstableList() []fs.DirEntry {
	f, err := os.Open(dirName)
	if err != nil {
		panic(err)
	}
	e, err := f.ReadDir(maxSstableSize)
	if err != nil {
		panic(err)
	}
	return e
}

// TODO: [portability] Depends on OS behavior that file order is ordered by desc.
func Search(searchKey string) (string, bool) {
	for _, s := range getSstableList() {

		f, err := os.Open(dirName + s.Name())
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
			// TODO: [perf] Skipping read value if the key matches.
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
	}

	return "", false
}
