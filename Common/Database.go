package Common

import (
	"buildinggit/DatabaseUtils"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"path"
)

type Database struct {
	DatabaseUtils.Backends
	path    string
	objects map[[20]byte]*Object
}

func newDatabase(path string) Database {
	return Database{path: path}
}

func (d Database) write(object Object) {
	contentBytes := []byte(object.content)
	content := object.kind + " " +
		string(rune(object.size)) + " " +
		base64.StdEncoding.EncodeToString(contentBytes)
	sum := sha1.Sum([]byte(content))
	d.writeToDisk(sum, object.content)
}

func (d Database) processWrite(tree databaseUtils.Tree) {

}

func store() {

}

func (d Database) HashObject(object []byte) [20]byte {
	return sha1.Sum(object)
}

func (d Database) writeToDisk(id [20]byte, content []byte) error {
	//var in bytes.Buffer
	//writer := zlib.NewWriter(&in)
	//writer.Write([]byte())
	//writer.Close()
	s := string(id[0]) + string(id[1])
	fileNameByte := id[2:]
	fileName := ""
	for b := range fileNameByte {
		fileName += string(b)
	}
	_, err2 := os.Stat(fileName)
	if err2 != nil {
		return errors.New("File already exist ")
	}
	err := os.Mkdir(path.Join(d.path, s), 777)

	open, err := os.Open(path.Join(d.path, s, fileName))
	if err != nil {
		return errors.New("File can not be opened ")
	}
	_, err = open.Write(content)
	if err != nil {
		return errors.New("File can not be write ")
	}
	open.Close()
	return nil
}

func readRawToBlob(path string) {

}

//func readObject(path string) {
//	var out bytes.Buffer
//	open, _ := os.Open(path)
//	r, _ := zlib.NewReader(open)
//	io.Copy(&out, r)
//
//}

func (db *Database) readObject(oid [20]byte) {
	db.LoadRaw(oid)
}

func generateTempName() {

}

func (db *Database) SerializeObject(blob databaseUtils.Blob) string {
	bytes := []byte(blob.ToS())
	return fmt.Sprintf("%s %s\000%s", blob.Type(), len(bytes), bytes)

}

func (db *Database) Store(blob databaseUtils.Blob) {
	object := db.SerializeObject(blob)
	blob.Oid = db.HashObject([]byte(object))
	db.WriteObject
}

func (db *Database) load(oid [20]byte) databaseUtils.Commit {
	//load files by object id
	if db.objects[oid] == nil {
		db.objects[oid] = db.readObject(oid)
	}
	return db.objects[oid]

}

func (db *Database) LoadTreeEntry(headOid [20]byte, targetPath string) *databaseUtils.Entry {
	commit := db.load(headOid)
	root := databaseUtils.NewEntry(commit.Tree, databaseUtils.TREE_MODE)
	return &databaseUtils.Entry{}
}

func LoadTreeList(oid []byte) {

}
