package index

//used to
import (
	"buildinggit"
	"buildinggit/repository"
	"encoding/json"
	"main/utils"
	"os"
	"strings"
	"time"
)

var REGULAR_MODE string = "100644"
var EXECUTABLE_MODE string = "100755"

const ENTRY_FORMAT string = "N10H40nZ*"

type Entry struct {
	name  string
	OId   []byte
	Stat  *main.FileInfo
	Path  string
	ctime *time.Time
	size  int64
	mtime *time.Time
}

func newEntry(name string, objectId []byte, stat *main.FileInfo) *Entry {

	return &Entry{name: name, OId: objectId, Stat: stat}
}

func UpdateStatus(stat repository.Status) {

}

func parseEntryFromBytes(bytes []byte) Entry {
	// deserialize the entries from []byte.
	s := string(bytes)
	splits := strings.Split(s, " ")
	var result os.FileInfo
	err := json.Unmarshal(bytes[24:], &result)
	info := main.ParseFileInfo(string(bytes[24:]))
	if err != nil {
		return Entry{name: splits[0], OId: []byte(splits[1]), Stat: info}
	}
	return Entry{}
}

func (entry Entry) mode() string {
	if entry.Stat.Mode > 111 {
		return EXECUTABLE_MODE
	}
	return REGULAR_MODE

}

func (entry Entry) toString() string {
	return entry.name + " __ " + string(entry.OId) + "__" + entry.Stat.Name + "\n"
}

func (entry Entry) ParentDirectories() []string {
	return utils.GetAncestors(entry.name)
}

func (entry Entry) baseName() string {
	split := strings.Split(entry.name, "/")
	return split[len(split)-1]
}

func (entry Entry) key() []byte {
	return entry.OId
}

func (entry Entry) StatMatch(stat os.FileInfo) bool {
	return entry.mode() == stat.Mode().String() && (entry.size == 0 || entry.size == stat.Size())
}

func parseEntry(entry string) *Entry {
	split := strings.Split(entry, "__")
	return newEntry(split[0], []byte(split[1]), main.ParseFileInfo(split[2]))

}
func (entry Entry) TimesMatch(stat os.FileInfo) bool {
	if entry.ctime == nil {
		return stat.ModTime().Equal(*entry.mtime)
	}
	return stat.ModTime().Equal(*entry.mtime)
}
