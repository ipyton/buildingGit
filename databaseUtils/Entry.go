package databaseUtils

// Entry for databaseUtils
type Entry struct {
	Oid  [20]byte
	Mode string
}

func NewEntry(oid string, mode string) *Entry {
	return &Entry{oid, mode}
}

func (e *Entry) isTree() bool {
	return e.Mode == "tree"
}
