package database

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
	"time"
)

type Commit struct {
	parents   []string
	tree      string
	author    Author
	committer Author
	message   string
}

func Parse(reader bufio.Reader) *Commit {
	headers := make(map[string][]string)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				// 输出最后一行未以换行符结尾的数据
				fmt.Print(line)
				break
			} else {
				fmt.Println("Error reading line:", err)
				break
			}
		}
		if line == "" {
			break
		}
		line = strings.TrimSpace(line)
		re := regexp.MustCompile(` +`)
		parts := re.Split(line, 2)
		headers[parts[0]] = append(headers[parts[0]], parts[1])
	}

	message, _ := reader.ReadString(0)

	return &Commit{committer: ParseAuthor(headers["author"][0]), author: ParseAuthor(headers["committer"][0]),
		parents: headers["parent"], tree: headers["tree"][0], message: message}
}

// merge needed?
func (commit *Commit) NeedMerge() bool {
	return len(commit.parents) > 1
}

func (commit *Commit) Parent() string {
	return commit.parents[0]
}

func (commit *Commit) Date() time.Time {
	return commit.committer.time
}

func (commit *Commit) TitleLine() string {
	return commit.message
}

func (commit *Commit) Type() string {
	return "commit"

}
func (commit *Commit) ToS() string {
	lines := []string{}
	lines = append(lines, fmt.Sprintf("tree {%s}", commit.tree))
	for i := range commit.parents {
		lines = append(lines, fmt.Sprintf("parent %s", commit.parents[i]))
	}
	lines = append(lines, fmt.Sprintf("author {%s}", commit.author))
	lines = append(lines, fmt.Sprintf("commmitter {%s}", commit.committer))
	lines = append(lines, "")
	result := strings.Join(lines, "\n")
	return result

}
