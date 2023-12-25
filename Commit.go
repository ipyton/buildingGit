package main

import (
	"strings"
	"time"
)

type Commit struct {
	treeId string
	author Author
	message string
	currentTime time.Time
	parent string
}

func newCommit(treeId string, author Author, message string, currentTime time.Time, parent string) *Commit {
	return &Commit{treeId: treeId, author: author, message: message, currentTime: currentTime, parent: parent}
}

func (commit Commit) toString() string {
	var lines []string
	lines = []string{commit.treeId, commit.author.toString(),commit.parent, "" ,commit.currentTime.String(),commit.message,}
	join := strings.Join(lines, "\n")
	return join
}
