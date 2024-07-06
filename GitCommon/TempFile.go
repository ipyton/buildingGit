package GitCommon

import (
	"math/rand"
	"os"
	path2 "path"
	"strconv"
)

type TempFile struct {
	dirName string
	path    string
	file    *os.File
}

func newTempFile(dirName, prefix string) *TempFile {
	return &TempFile{dirName: dirName, path: path2.Join(dirName, prefix+"_"+strconv.Itoa(rand.Int())), file: nil}
}

func (tmp *TempFile) write(data string) {
	if tmp.file == nil {
		tmp.file, _ = os.Open(tmp.path)
	}
	tmp.file.WriteString(data)
}

func (tmp *TempFile) move(name string) {
	tmp.file.Close()
	os.Rename(tmp.path, path2.Join(tmp.dirName, name))
}
func (tmp *TempFile) OpenFile() {
	tmp.file, _ = os.OpenFile(tmp.path, os.O_RDWR|os.O_CREATE|os.O_EXCL, os.FileMode(0644))
	os.Mkdir(tmp.dirName, os.FileMode(0755))

}
