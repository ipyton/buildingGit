package main

import (
	"fmt"
	"io/fs"
	"os"
	path2 "path"
	"time"
)

func init(){

}

func commandDispatcher(command string, args []string) {
	path, _ := os.Getwd()
	if command == "init" {
		handleInit(path)
	}
	if command == "commit" {
		handleCommit(path)
	}
	if command == "add" {
		handleAdd(path, args)
	}
}

func handleAdd(path string, args []string) {
	if len(args) != 1 {
		fmt.Println("needs an argument!")
	}
	gitPath := path2.Join(path,".git")
	// objectsPath := path2.Join(gitPath, "objects")
	//workspace := Workspace{path: path,
	//	ignore: []string{".", "..", ".git"}}
	database := newDatabase(path2.Join(gitPath, "objects"))
	index := newIndex(path2.Join(gitPath, "index"))
	for _, argPath := range args {
		pathInDirectory := path2.Join(path, argPath)
		byPath := listFilesByPath(pathInDirectory)
		for _, file := range byPath {
			data := readFile(file)
			stat := statFile(file)
			blob := newObject(data, "blob")
			database.write(blob)
			index.add(path, blob.id, stat)
		}
	}
	index.writeUpdates()

}

func handleInit(path string) {
	path = path2.Join(path, ".git")
	err := os.Mkdir(path, os.FileMode(777))
	if err != nil {
		return
	}
	var folders  = [...]string{"objects","refs"}
	for _,v :=range folders {
		fmt.Println(v)
		nPath := path2.Join(path, v)
		fmt.Println(nPath)
		err := os.Mkdir(nPath, fs.FileMode(777))
		if err != nil {
			fmt.Println("has already exist")
			return
		}
	}
}

func handleCommit(path string){
	entries :=make([]Entry,0)
	gitPath := path2.Join(path,".git")
	objectsPath := path2.Join(gitPath, "objects")
	fmt.Println(objectsPath)
	workspace := Workspace{path: path,
		ignore: []string{".", "..", ".git"}}
	files := workspace.listFiles()
	database := newDatabase(objectsPath)
	for _, filePath := range files {
		data := readFile(filePath)
		object := newObject(data,"blob")
		database.write(object)
		entries = append(entries, Entry{name: filePath, objectId: object.id})
	}

	tree := newTree(entries)
	object := Object{kind: "tree", content: []byte(tree.toString())}
	database.write(object)
	name := os.Getenv("GIT_AUTHOR_NAME")
	email := os.Getenv("GIT_AUTHOR_EMAIL")
	author := newAuthor(name, email)
	var message string
	fmt.Scanln(&message)

	ref := newRef(gitPath)
	parent := ref.readHead()

	commit := newCommit(object.id,author, message,time.Now(), parent)
	commitObject := Object{kind: "commit", content: []byte(commit.toString())}
	database.write(commitObject)
	headPath := path2.Join(gitPath, "HEAD")
	file, _ := os.OpenFile(headPath, os.O_CREATE|os.O_WRONLY, 0777)
	file.Write([]byte("(root-commit) " + commitObject.id + commit.treeId))
}




func main(){
	//commandDispatcher("commit")
	//writeToDisk()


}