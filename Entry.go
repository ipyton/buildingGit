package main

import (
	"encoding/json"
	"os"
	"strings"
)

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

func parseEntryFromBytes(bytes []byte) Entry{
	s := string(bytes)
	splits := strings.Split(s, " ")
	var result os.FileInfo
	err := json.Unmarshal(bytes[24:], &result)
	if err != nil {
		return Entry{name: splits[0], objectId: splits[1], stat: result}
	}
	return Entry{}
}

func (entry Entry) mode() string {
	if entry.stat.Mode() > 111 {
		return EXECUTABLE_MODE
	}
	return REGULAR_MODE

}

func (entry Entry) toString() string {
	return entry.name + " __ " + entry.objectId + entry.stat.Name()
}

func (entry Entry) parentDirectories() []string{

}

