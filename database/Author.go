package database

import (
	"fmt"
	"strings"
	"time"
)

type Author struct {
	Name  string
	Email string
	time  time.Time
}

func (Author) parse(content string) Author {
	splited := strings.Split(content, "")
	name, email, mod_time := splited[0], splited[1], splited[2]
	t, _ := time.Parse("", mod_time)
	return Author{
		name,
		email, t,
	}

}

func (a Author) shortDate() string {
	return a.time.Format("yyyy-MM-dd")
}

func (a Author) readableTime() string {
	return a.time.Format("yyyy-MM-dd HH:mm:ss")
}
func (a Author) toS() string {
	timestamp := a.time.Format("yyyy-MM-dd HH:mm:ss")
	return fmt.Sprintf("%s <%s> %s", a.Name, a.Email, timestamp)
}
