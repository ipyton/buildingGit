package databaseUtils

import (
	"bufio"
	index2 "buildinggit/indexUtils"
	"sort"
	"strings"
)

type Tree struct {
	ENTRY_FORMAT string
	objectType   string
	objects      []*index2.Entry
	mode         string
}

func (tree Tree) getObjectType(objectType string) string {
	return "tree"
}

func (tree Tree) ToString() string {
	tree.sortByName()
	result := ""
	for _, item := range tree.objects {
		result += item.mode() + item.name + item.objectId
	}
	return result
}

func (tree Tree) sortByName() {
	sort.SliceStable(tree.objects, func(a int, b int) bool {
		return tree.objects[a].name < tree.objects[b].name
	})

}

func (tree Tree) parseTree(scanner bufio.Scanner) string {
	// read from file and go back
	entries := make(map[string]string)

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")

		entries[split[1]] = index2.newEntry(entries[0], entries[2])
	}

}

func newTree(entries []*index2.Entry) *Tree {

	return &Tree{objectType: "tree", mode: "100644", objects: entries}
}

func parseTree(entryString string) *Tree {
	split := strings.Split(entryString, "\n")
	entries := make([]*index2.Entry, 10, 10)
	for _, value := range split {
		entry := index2.parseEntry(value)
		entries = append(entries, entry)
	}
	return newTree(entries)
}

func (tree Tree) traverse() {

}
