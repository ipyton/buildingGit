package GitCommon

import (
	"crypto/sha1"
	"fmt"
	"hash"
	"os"
)

type SumChecker struct {
	file         *os.File
	sha1Digest   hash.Hash
	checkSumSize int
}

// initialize
func newSumChecker(path string) SumChecker {
	fileToCheck, error := os.Open(path)
	if error != nil {
		fmt.Println("file does not exist!")
	}
	return SumChecker{file: fileToCheck, sha1Digest: sha1.New()}
}

func (checker SumChecker) read(size int) bool {
	buffer := make([]byte, size, size)
	read, err := checker.file.Read(buffer)
	if read != size || err != nil {
		return false
	}
	_, err = checker.sha1Digest.Write(buffer)
	if err != nil {
		return false
	}
	return true
}

func (checker SumChecker) verify() bool {
	buffer := make([]byte, checker.checkSumSize, checker.checkSumSize)
	numberOfBytes, error := checker.file.Read(buffer)
	if numberOfBytes != checker.checkSumSize || error != nil {
		return false
	}
	sumResult := checker.sha1Digest.Sum(nil)
	if len(buffer) != len(sumResult) {
		return false
	}
	for index, content := range buffer {
		if content != sumResult[index] {
			return false
		}
	}
	return true
}
