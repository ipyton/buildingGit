package Common

import (
	"errors"
	"os"
	"path"
	"strings"
)

type Refs struct {
	pathname string
}

// references of commits in repositoryUtils.
// and git stash

func newRef(pathName string) Refs {
	return Refs{pathname: pathName}
}

func (ref Refs) updateHead(objectId string) error {
	lock := main.newLock(ref.getHeadPath())
	if !lock.lock() {
		return errors.New("lock could not be acquired exist")
	}

	//flag := os.O_WRONLY | os.O_CREATE
	//file, _ := os.OpenFile(ref.getHeadPath(), flag, 0777)
	//file.WriteString(objectId)
	lock.write(objectId)
	lock.write("\\n")
	lock.commit()
	return nil
}

func (ref Refs) readHead() string {
	headPath := ref.getHeadPath()
	_, err := os.Stat(headPath)

	if !errors.Is(err, os.ErrExist) {
		file, _ := os.ReadFile(headPath)

		s := string(file)
		trim := strings.Trim(s, " ")
		return trim
	}
	return ""
}

func (ref Refs) getHeadPath() string {
	return path.Join(ref.pathname, "HEAD")
}
