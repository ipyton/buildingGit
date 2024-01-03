package main

import (
	"crypto/sha1"
	"github.com/wangjia184/sortedset"
	"hash"
	"os"
	"strconv"
	"strings"
)


type Index struct {
	file * os.File
	path string
	lock Lock
	entries map[string] Entry
	keys * sortedset.SortedSet
	sha1Digest    hash.Hash
	HeaderSize   int
	HeaderFormat string
	Signature    string
	Version int
	changed bool
	EntryMinSize int
	EntryBlock int
}


func newIndex(path string) Index {
	open, _ := os.Open(path)
	//stringList := []int{1,}
	//keys := [..]StringHeap{"1"}
	set := sortedset.New()
	//heap.Init(keys)

	return Index{file: open, lock: newLock(path), path: path, sha1Digest: sha1.New(), HeaderSize: 12,
		HeaderFormat: "a4N2", Signature: "DIRC", Version: 2, keys: set}
}

func (this Index) clear() bool {
	for key, _ := range this.entries {
		delete(this.entries, key)
	}
	this.keys = sortedset.New()
	this.sha1Digest.Reset()
	return true
}


func (this Index) add(pathName string, oid string, state os.FileInfo) {
	entry := newEntry(pathName, oid, state)
	this.entries[pathName] = entry
	this.changed = true

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
		headerSize := this.readHeader(*this.file)
		this.readEntries(*this.file, headerSize)
		checker := newSumChecker(this.path)
		checker.verify()

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
	if this.changed {
		return this.lock.rollback()
	}

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



func (this Index) readHeader(file os.File) int {
	buffer := make([]byte, this.HeaderSize)
	read, err := file.Read(buffer)
	if err != nil || read != this.HeaderSize {
		return -1
	}
	split := strings.Split(string(buffer), "@")
	signature, versionString, countString := split[0],split[1], split[2]
	version, err := strconv.Atoi(versionString)
	if signature != this.Signature || version != this.Version {
		return -1
	}
	count, err := strconv.Atoi(countString)
	if err != nil {
		return -1
	}
	return count
}

//this one should be changed
func (this Index) readEntries(file os.File, count int) bool {
	for i := 0; i < count; i ++ {
		buffer := make([]byte, this.EntryMinSize, this.EntryMinSize)
		size, err := file.Read(buffer)
		for {
			if buffer[size - 1] != '\n' {
				appendBuffer := make([]byte, 8, 8)
				size, err := file.Read(appendBuffer)
				if size != 0 || err != nil {
					return false
				}
				buffer = append(buffer,appendBuffer...)
			} else {
				break
			}
		}
		if  err != nil {
			return false
		}
		entry := parseEntryFromBytes(buffer)
		this.storeEntry(entry)
	}
	return true

}

func (this Index) storeEntry(entry Entry) {
	this.keys.AddOrUpdate(entry.key(),20 ,nil)
	// this.keys = append(this.keys, entry.key)
	this.entries[entry.key()] = entry

}

func (this Index) opEachEntry(op func(entry Entry)) {
	for _,v := range this.entries {
		op(v)
	}
}

func (this Index) discardConflicts(entry Entry){
	for _, directory := range entry.parentDirectories() {
		this.keys.Remove(entry.key())
		delete(this.entries, directory)
	}
}

func basename() {

}

func rollBack() {

}

func write() {

}

func writeCheckSum() {

}
