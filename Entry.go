package main

import (
	"encoding/json"
	"main/utils"
	"os"
	"strings"
)

var REGULAR_MODE string = "100644"
var EXECUTABLE_MODE string = "100755"

type Entry struct {
	name string
	objectId string
	stat * FileInfo
	path string
}


func newEntry(name string, objectId string, stat * FileInfo) *Entry {

	return & Entry{name: name,objectId: objectId, stat: stat}
}

func parseEntryFromBytes(bytes []byte) Entry{
	s := string(bytes)
	splits := strings.Split(s, " ")
	var result os.FileInfo
	err := json.Unmarshal(bytes[24:], &result)
	info := parseFileInfo(string(bytes[24:]))
	if err != nil {
		return Entry{name: splits[0], objectId: splits[1], stat: info}
	}
	return Entry{}
}

func (entry Entry) mode() string {
	if entry.stat.mode > 111 {
		return EXECUTABLE_MODE
	}
	return REGULAR_MODE

}

func (entry Entry) toString() string {
	return entry.name + " __ " + entry.objectId +  "__" + entry.stat.name + "\n"
}

func (entry Entry) parentDirectories() []string{
	return utils.GetAncestors(entry.name)
}

func (entry Entry) baseName() string {
	split := strings.Split(entry.name, "/")
	return split[len(split) - 1]
}


func (entry Entry) key() string {
	return entry.objectId
}

func parseEntry(entry string) * Entry {
	split := strings.Split(entry, "__")
	return newEntry(split[0], split[1], parseFileInfo(split[2]))

}
