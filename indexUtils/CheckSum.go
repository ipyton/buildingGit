package indexUtils

import (
	bytes2 "bytes"
	"hash"
	"os"
)

const CHECK_SUM_SIZE = 20

type Checksum struct {
	file os.File
	hash hash.Hash
}

func (checksum *Checksum) write(data string) {
	checksum.file.Write([]byte(data))
	checksum.hash.Write([]byte(data))

}

func (checksum *Checksum) writeChecksum() {
	checksum.file.Write(checksum.hash.Sum(nil))
}

func (checksum *Checksum) read(size int64) string {
	bytes := make([]byte, size)

	length, _ := checksum.file.Read(bytes)
	if int64(length) != size {
		return "error"
	}
	checksum.hash.Write(bytes)
	return string(bytes)

}

func (checksum *Checksum) verifyChecksum() bool {
	bytes := make([]byte, CHECK_SUM_SIZE)
	n, err := checksum.file.Read(bytes)
	if err != nil && n != CHECK_SUM_SIZE {
		return false
	}
	if bytes2.Compare(bytes, checksum.hash.Sum(nil)) == 0 {
		return true
	}
	return false

}
