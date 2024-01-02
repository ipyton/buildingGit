package main

import (
	"crypto/sha1"
	"hash"
	"os"
)

type StringHeap []string



type Index struct {
	file * os.File
	path string
	lock Lock
	entries map[string] Entry
	keys StringHeap
	sha1Digest hash.Hash
}


func newIndex(path string) Index {
	open, _ := os.Open(path)
	//stringList := []int{1,}
	//keys := [..]StringHeap{"1"}

	//heap.Init(keys)

	return Index{file: open, lock: newLock(path), path: path, sha1Digest: sha1.New()}
}

func (this Index) clear() bool {
	for key, _ := range this.entries {
		delete(this.entries, key)
	}

	for


}


func (this Index) add(pathName string, oid string, state os.FileInfo) {
	entry := newEntry(pathName, oid, state)
	this.entries[pathName] = entry
}


//lock for update.
func (this Index) loadForUpdate() bool {
	if this.lock.lock() {
		this.load(this.path)
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


func (this Index) write(data string)  {
	this.lock.write(data)
	this.sha1Digest.Write([]byte(data))
}


func (this Index) finishWrite() {
	this.lock.write(string(this.sha1Digest.Sum(nil)))
	this.lock.commit()

}


func (Index) getCheckSumLength() int {
	return 20
}

func (this Index) read(size int) []byte{
	var result []byte
	this.file.Read(result)
	return result
}


//func (this Index) checkSum() bool {
//	read := this.read(this.getCheckSumLength())
//
//	return true
//}


func (this Index) writeUpdates() bool {
	if !this.lock.lock() {
		return false
	}
	header := "DIRC" + string(2) + string(len(this.entries))
	this.write(header)
	for _, entry := range this.entries {
		this.write(entry.toString())
	}
	this.finishWrite()
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
