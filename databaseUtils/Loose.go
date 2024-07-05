package DatabaseUtils

import (
	"bufio"
	bytes2 "bytes"
	"compress/zlib"
	"os"
	"path"
	"strconv"
	"strings"
)

// a loose file stands for not compressed files in the project.
// packed file is used to compress files

type Loose struct {
	pathname string
}

func (loose *Loose) Has(oid [20]byte) bool {
	info, _ := os.Stat(path.Join(loose.pathname, string(oid[:])))

	// info := os.FileInfo(path.Join(loose.pathname, loose.ObjectPath(oid)))
	return !info.IsDir()

}

func (loose *Loose) LoadInfo(oid string) *Raw {
	path := loose.ObjectPath(oid)
	reader, err := os.Open(path)
	bytes := make([]byte, 128)
	n, err := reader.Read(bytes)
	if n != 128 || err != nil {
		println("unsupported ")
	}
	r := bytes2.NewReader(bytes)
	newReader, err := zlib.NewReader(r)
	newReader.Read(bytes)
	r2 := bufio.NewReader(newReader)
	Type, err := r2.ReadString(' ')
	if err != nil {
		println("error reading ")
	}
	readString, err := r2.ReadString('\000')
	if err != nil {
		size, _ := strconv.Atoi(readString[0 : len(readString)-2])
		return &Raw{Data: reader, Type: Type, Size: size}
	}
	return nil
}

func (loose *Loose) LoadRaw(oid string) *Raw {
	return loose.LoadInfo(oid)
}

func (loose *Loose) PrefixMatch(name string) []string {

	objectPath := loose.ObjectPath(name)
	dirs, err := os.ReadDir(path.Dir(objectPath))
	oids := make([]string, len(dirs))
	if err != nil {
		for i := range dirs {
			entry := dirs[i]
			name := entry.Name()
			oids = append(oids, name)
		}
	}
	result := make([]string, len(oids))
	for i := range oids {

		if strings.HasPrefix(oids[i], name) {
			result = append(result, oids[i])
		}
	}
	return result

}
func (loose *Loose) WriteObject(oid string, content string) {
	objectPath := loose.ObjectPath(oid)
	_, err := os.Open(objectPath)
	if err != nil {
		if os.IsNotExist(err) {
			os.CreateTemp()

		}
	}
	return
}

func (loose *Loose) ObjectPath(oid string) string {
	return path.Join(string(oid[0:2]), string(oid[2:]))
}
