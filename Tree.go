package main

import (
	"sort"
	"strings"
)

type Tree struct {
	ENTRY_FORMAT string
	objectType   string
	objects [] * Entry
	mode string
}

func (tree Tree) getObjectType(objectType string) string {
	return "tree"
}

func (tree Tree) toString() string {
	tree.sortByName()
	result := ""
	for _, item := range tree.objects {
		result += item.mode() + item.name + item.objectId
	}
	return result
}

func (tree Tree) sortByName(){
		sort.SliceStable(tree.objects, func(a int, b int) bool {
		return tree.objects[a].name < tree.objects[b].name
	})

}

func newTree(entries [] * Entry) *Tree {
	return &Tree{objectType:"tree", mode: "100644", objects: entries}
}

func parseTree(entryString string) * Tree{
	split := strings.Split(entryString, "\n")
	entries := make([] * Entry, 10, 10)
	for _, value := range split {
		entry := parseEntry(value)
		entries = append(entries, entry)
	}
	return newTree(entries)
}


func (tree Tree) traverse() {

}
