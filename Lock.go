package main

import (
	"os"
	path2 "path"
)

type Lock struct {
	filePath string
	lockPath string
	status bool
	file *os.File
}

func newLock(path string) Lock {
	return Lock{filePath: path, lockPath: path2.Ext(".lock")}
}

func (lock Lock) lock() bool {
	if !lock.status {
		flags := os.O_RDWR | os.O_CREATE | os.O_EXCL
		file, _ := os.OpenFile(lock.filePath, flags, 0777)
		lock.file = file
		return true
	}
	return false

}

func (lock Lock) write(content string) {
	lock.raiseOnStaleLock()
	_, err := lock.file.WriteString(content)
	if err != nil {
		return
	}
}

func (lock Lock) commit() {
	lock.raiseOnStaleLock()
	lock.status = false
	err := os.Rename(lock.lockPath, lock.filePath)
	if err != nil {
		return
	}
	lock.file = nil
}

func (lock Lock) raiseOnStaleLock() bool {
	if !lock.status {
		return false
	}
	return true
}