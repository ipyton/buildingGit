package GitCommon

import "strings"

type Author struct {
	name  string
	email string
}

func newAuthor(name string, email string) Author {
	return Author{name: name, email: email}
}

func (author Author) toString() string {
	return author.name + " " + author.email
}

func parseAuthor(author string) Author {
	split := strings.Split(author, " ")
	return newAuthor(split[0], split[1])
}
