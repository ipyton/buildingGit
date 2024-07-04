package databaseUtils

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
	"time"
)

type Commit struct {
	Parents   []string
	Tree      string
	Author    Author
	Committer Author
	Message   string
}

func ParseCommit(reader bufio.Reader) *Commit {
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

	return &Commit{Committer: ParseAuthor(headers["author"][0]), Author: ParseAuthor(headers["committer"][0]),
		Parents: headers["parent"], Tree: headers["tree"][0], Message: message}
}

// merge needed?
func (commit *Commit) NeedMerge() bool {
	return len(commit.Parents) > 1
}

func (commit *Commit) Parent() string {
	return commit.Parents[0]
}

func (commit *Commit) Date() time.Time {
	return commit.Committer.time
}

func (commit *Commit) TitleLine() string {
	return commit.Message
}

func (commit *Commit) Type() string {
	return "commit"

}
func (commit *Commit) ToS() string {
	lines := []string{}
	lines = append(lines, fmt.Sprintf("tree {%s}", commit.Tree))
	for i := range commit.Parents {
		lines = append(lines, fmt.Sprintf("parent %s", commit.Parents[i]))
	}
	lines = append(lines, fmt.Sprintf("author {%s}", commit.Author))
	lines = append(lines, fmt.Sprintf("commmitter {%s}", commit.Committer))
	lines = append(lines, "")
	result := strings.Join(lines, "\n")
	return result
}
