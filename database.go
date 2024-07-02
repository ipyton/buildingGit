package main

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"io"
	"os"
	"path"
)

type Database struct {
	path string
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

func (d Database) processWrite(tree Tree) {

}

func store() {

}

func hash_object() {

}

func serialize_object() {

}
func hash_content() {

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

func readObject(path string) {
	var out bytes.Buffer
	open, _ := os.Open(path)
	r, _ := zlib.NewReader(open)
	io.Copy(&out, r)

}

func generateTempName() {

}
