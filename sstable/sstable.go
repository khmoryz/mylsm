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

const DirName = "/tmp/mylsm/"
const suffix = "_sstable.mylsm"
const maxSstableSize = 1000
const compactionThreshold = 10
const blockSize = 1 * 1024

var blockSeparator = []byte{0, 0, 0, 0}

// File format of SSTable
// ↓offset=bockSize*0 (indexed)                      ↓offset=bockSize*1 (indexed)
// ┌────────────┬─────┬──────────────┬───────┬─────┬────────────┬─────┬──────────────┬───────┬───────────┬────────────┬─────┬────────┬────────────┬─────┬───────────┬──────────────┐
// │ Key Length │ Key │ Value Length │ Value │ ... │ Key Length │ Key │ Value Length │ Value │ Separator │ Key Length │ Key │ Offset │ Key Length │ ... │ Separator │ Index Offset │
// └────────────┴─────┴──────────────┴───────┴─────┴────────────┴─────┴──────────────┴───────┴───────────┴────────────┴─────┴────────┴────────────┴─────┴───────────┴──────────────┘
// └────────────────────────────────────── Data Block ───────────────────────────────────────┘           └─────────────── Partial Index ────────────────┘           └── Metadata ──┘

var Memt Table

type Table struct {
	Kvs []Kv
}

type Kv struct {
	Key       string
	Value     string
	TombStone bool
}

type entry struct {
	key    string
	offset int64
}

func genFileName() string {
	now := time.Now().UnixNano()
	return strconv.FormatInt(now, 10) + suffix
}

func Flush(t *Table) error {
	if err := Write(t); err != nil {
		return err
	}

	// compaction
	if len(getSstableList()) >= compactionThreshold {
		Compaction()
	}
	// initialization
	Memt = Table{}
	return nil
}

func isBlockHead(cur int64, curBlock int, size int) bool {
	return cur+int64(size) >= int64(blockSize*curBlock)
}

func calcByteSize(k Kv) int {
	return 4 + len(k.Key) + 4 + len(k.Value)
}

func Write(t *Table) error {
	// uniq
	uniq := make(map[string]Kv, 0)
	for _, k := range t.Kvs {
		uniq[k.Key] = Kv{Key: k.Key, Value: k.Value, TombStone: k.TombStone}
	}

	// sort
	var sst Table
	for _, v := range uniq {
		sst.Kvs = append(sst.Kvs, Kv{Key: v.Key, Value: v.Value, TombStone: v.TombStone})
	}
	sort.Slice(sst.Kvs, func(i, j int) bool { return sst.Kvs[i].Key < sst.Kvs[j].Key })

	// write
	if err := os.MkdirAll(DirName, 0755); err != nil {
		panic(err)
	}

	f, err := os.OpenFile(DirName+genFileName(), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	partialIndex := []entry{}

	// write datablock
	for _, v := range sst.Kvs {

		// collect index
		cur, _ := f.Seek(0, io.SeekCurrent)
		if isBlockHead(cur, len(partialIndex), calcByteSize(v)) {
			partialIndex = append(partialIndex, entry{key: v.Key, offset: cur})
		}

		binary.Write(f, binary.LittleEndian, int32(len(v.Key)))
		f.Write([]byte(v.Key))
		binary.Write(f, binary.LittleEndian, int32(len(v.Value)))
		f.Write([]byte(v.Value))
		binary.Write(f, binary.LittleEndian, v.TombStone)
	}

	// write separator
	f.Write(blockSeparator)

	// write index block
	idxOfs, _ := f.Seek(0, io.SeekCurrent)
	for _, v := range partialIndex {
		binary.Write(f, binary.LittleEndian, int32(len(v.key)))
		f.Write([]byte(v.key))
		binary.Write(f, binary.LittleEndian, v.offset)
	}

	// write separator
	f.Write(blockSeparator)

	// write metadata block
	binary.Write(f, binary.LittleEndian, int32(idxOfs))

	return nil
}

func getSstableList() []fs.DirEntry {
	f, err := os.Open(DirName)
	if err != nil {
		panic(err)
	}
	e, err := f.ReadDir(maxSstableSize)
	if err != nil {
		panic(err)
	}
	return e
}

// func getIndex(f *os.File) []entry {

// }

// TODO: [portability] Depends on OS behavior that file order is ordered by desc.
func Search(searchKey string) (string, bool) {
	for _, s := range getSstableList() {

		f, err := os.Open(DirName + s.Name())
		if err != nil {
			panic(err)
		}
		defer f.Close()

		for {
			// idx:=getIndex()

			// Read key
			var klen int32
			if err := binary.Read(f, binary.LittleEndian, &klen); err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}

			// end data block
			if klen == 0 {
				break
			}

			k := make([]byte, klen)
			if _, err := f.Read(k); err != nil {
				panic(err)
			}

			// Read value
			// TODO: [perf] Skipping read value if the key matches.
			var vlen int32
			if err := binary.Read(f, binary.LittleEndian, &vlen); err != nil {
				panic(err)
			}
			v := make([]byte, vlen)
			if _, err := f.Read(v); err != nil {
				panic(err)
			}

			// Read tombstone
			var ts bool
			if err := binary.Read(f, binary.LittleEndian, &ts); err != nil {
				panic(err)
			}

			// compare
			if string(k) == searchKey {
				return string(v), !ts
			}
		}
	}
	return "", false
}

func ReadRow(f *os.File) (Kv, error) {
	// Read key
	var klen int32
	if err := binary.Read(f, binary.LittleEndian, &klen); err != nil {
		return Kv{}, err
	}
	k := make([]byte, klen)
	if _, err := f.Read(k); err != nil {
		panic(err)
	}

	// Read value
	// TODO: [perf] Skipping read value if the key matches.
	var vlen int32
	if err := binary.Read(f, binary.LittleEndian, &vlen); err != nil {
		return Kv{}, err
	}
	v := make([]byte, vlen)
	if _, err := f.Read(v); err != nil {
		return Kv{}, err
	}

	// Read tombstone
	var ts bool
	if err := binary.Read(f, binary.LittleEndian, &ts); err != nil {
		return Kv{}, err
	}

	return Kv{Key: string(k), Value: string(v), TombStone: ts}, nil
}

func ReadTable(f *os.File, t *Table) {
	for {
		kv, err := ReadRow(f)
		if err == io.EOF {
			break
		}
		t.Kvs = append(t.Kvs, kv)
	}
}

// Compact sstable into one file.
func Compaction() error {
	var ct Table
	s := getSstableList()

	// read sstable order by asc.
	for i := len(s) - 1; i >= 0; i-- {
		f, err := os.Open(DirName + s[i].Name())
		if err != nil {
			panic(err)
		}
		ReadTable(f, &ct)
		f.Close()
	}

	Write(&ct)

	// delete compactioned sstable.
	for _, t := range s {
		if err := os.Remove(DirName + t.Name()); err != nil {
			panic(err)
		}
	}
	return nil
}
