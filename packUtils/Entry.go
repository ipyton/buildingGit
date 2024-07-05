package packUtils

import (
	"os"
	path2 "path"
)

type Entry struct {
	Oid   string
	Info  os.FileInfo
	Path  string
	Ofs   bool
	Delta string
	Depth int
}

func initializeEntry(oid string, info os.FileInfo, path string, ofs bool) *Entry {
	return &Entry{Oid: oid, Info: info, Path: path, Ofs: ofs}
}

func (entry *Entry) sortKey() []string {
	return []string{packedType(), path2.Base(entry.Path), path2.Dir(entry.Path), os.FileInfo()}
}

func assignDelta() string {

}

func packedType() string {

}

func packedSize() {

}

func deltaPrefix() {

}
