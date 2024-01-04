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
	return & Commit{treeId: treeId, author: author, message: message, currentTime: currentTime, parent: parent}
}

func (commit Commit) toString() string {
	var lines []string
	lines = []string{commit.treeId, commit.author.toString(),commit.parent, "" ,commit.currentTime.String(),commit.message,}
	join := strings.Join(lines, "\n")
	return join
}

func parseCommit(commitMessage string) * Commit {
	split := strings.Split(commitMessage, "\n")
	parse, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", split[4])
	if err != nil {
		return nil
	}
	return newCommit(split[0], parseAuthor(split[1]), split[5], parse, split[2])
}

