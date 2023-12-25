package main

type Entry struct {
	name string
	objectId string
}

func newEntry(name string, objectId string) Entry {
	return Entry{name: name,objectId: objectId}
}