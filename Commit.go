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
}

func newCommit(treeId string, author Author, message string, currentTime time.Time) *Commit {
	return &Commit{treeId: treeId, author: author, message: message, currentTime: currentTime}
}

func (commit Commit) toString() string {
	var lines []string
	lines = []string{commit.treeId, commit.author.toString(), commit.message, commit.currentTime.String(),}
	join := strings.Join(lines, "\n")
	return join
}
