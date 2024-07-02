package database

type Entry struct {
	oid  string
	mode string
}

func NewEntry(oid string, mode string) *Entry {
	return &Entry{oid, mode}
}

func (e *Entry) isTree() bool {
	return e.mode == "tree"
}
