package main

import "os"

var REGULAR_MODE string = "100644"
var EXECUTABLE_MODE string = "100755"

type Entry struct {
	name string
	objectId string
	stat os.FileInfo
}

func newEntry(name string, objectId string, stat os.FileInfo) Entry {
	return Entry{name: name,objectId: objectId, stat: stat}
}

func (entry Entry) mode() string{
	if entry.stat.Mode() > 111 {
		return EXECUTABLE_MODE
	}
	return REGULAR_MODE

}

