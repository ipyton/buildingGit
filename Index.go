package main

import (
	"container/heap"
	"fmt"
	"os"
)

type StringHeap []string



type Index struct {
	file * os.File
	path string
	lock Lock
	entries map[string] Entry
	keys StringHeap
}


func newIndex(path string) Index {
	open, _ := os.Open(path)
	stringList := []int{1, }
	keys :=[..]StringHeap{"1"}
	heap.Init(keys)

	return Index{file: open, lock: newLock("./lock"), path: path}
}


func (this Index) add(pathName string, oid string, state os.FileInfo) {
	entry := newEntry(pathName, oid, state)
	entry.ke
	this.key
}

func (this Index) loadForUpdate() bool {
	if this.lock.lock(){
		this.load()
		return true
	}
	return false
}

func (this Index) openIndexFile(path string) bool {

	indexFile, openError := os.OpenFile(path, os.O_RDONLY, 777)
	if openError != nil{
		this.file = indexFile
		return true
	}
	return false

}

func (this Index) load(path string) {
	this.openIndexFile(path)
	if this.file != nil {

	}

}

func (Index) getCheckSumLength() int {
	return 20
}

func (this Index) read(size int) []byte{
	var result []byte
	this.file.Read(result)
	return result
}


func (this Index) checkSum() bool {
	read:= this.read(this.getCheckSumLength())

	return true
}

func index() {

}

func readHeader() {

}

func readEntries() {

}

func storeEntry() {

}

func add(path string, objectId string, status string){


}

func writeUpdates() {

}

func rollBack() {

}

func write() {

}

func writeCheckSum() {

}
