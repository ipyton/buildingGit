package main

// This file is used to get the command from the console and dispatch to the functions.
import (
	database2 "buildinggit/DatabaseUtils"
	"buildinggit/GitCommon"
	index2 "buildinggit/indexUtils"
	"fmt"
	"io/fs"
	"os"
	path2 "path"
	"time"
)

func init() {

}

func commandDispatcher(command string, args []string) {
	path, _ := os.Getwd()

	if command == "add" {
		handleAdd(path, args)
	}

	if command == "handle" {
		handleStatus()
	}

	if command == "branch" {

	}

	if command == "checkout" {

	}

	if command == "cherry_pick" {

	}

	if command == "commit" {
		handleCommit(path)
	}

	if command == "config" {

	}

	if command == "diffUtils" {

	}

	if command == "fetch" {

	}

	if command == "init" {
		handleInit(path)
	}

	if command == "log" {

	}

	if command == "merge" {

	}

	if command == "push" {

	}

	if command == "receive_pack" {

	}

	if command == "rev_list" {

	}

	if command == "revert" {

	}

	if command == "rm" {

	}

	if command == "status" {

	}

	if command == "upload_pack" {

	}

}

func handleAdd(path string, args []string) {
	if len(args) != 1 {
		fmt.Println("needs an argument!")
	}
	gitPath := path2.Join(path, ".git")
	// objectsPath := path2.Join(gitPath, "objects")
	//workspace := Workspace{path: path,
	//	ignore: []string{".", "..", ".git"}}
	database := newDatabase(path2.Join(gitPath, "objects"))
	index := newIndex(path2.Join(gitPath, "indexUtils"))
	index.loadForUpdate()
	for _, argPath := range args {
		pathInDirectory := path2.Join(path, argPath)
		byPath := listFilesByPath(pathInDirectory)
		for _, file := range byPath {
			data := readFile(file)
			stat := statFile(file)
			blob := GitCommon.newObject(data, "blob")
			database.write(blob)
			index.add(path, blob.id, stat)
		}
	}
	index.writeUpdates()

}

func handleInit(path string) {
	// create paths needed
	path = path2.Join(path, ".git")
	err := os.Mkdir(path, os.FileMode(777))
	if err != nil {
		return
	}
	var folders = [...]string{"objects", "refs"}
	for _, v := range folders {
		fmt.Println(v)
		nPath := path2.Join(path, v)
		fmt.Println(nPath)
		err := os.Mkdir(nPath, fs.FileMode(777))
		if err != nil {
			fmt.Println("has already exist")
			return
		}
	}
	// create configuration file

	// refs

}

func handleCommit(path string) {
	entries := make([]*index2.Entry, 0)
	gitPath := path2.Join(path, ".git")
	objectsPath := path2.Join(gitPath, "objects")
	fmt.Println(objectsPath)
	workspace := GitCommon.Workspace{path: path,
		ignore: []string{".", "..", ".git"}}
	files := workspace.listFiles()
	database := newDatabase(objectsPath)
	for _, filePath := range files {
		data := readFile(filePath)
		object := GitCommon.newObject(data, "blob")
		database.write(object)
		entries = append(entries, &index2.Entry{name: filePath, objectId: object.id})
	}

	tree := database2.newTree(entries)
	object := GitCommon.Object{kind: "tree", content: []byte(tree.toString())}
	database.write(object)
	name := os.Getenv("JIT_AUTHOR_NAME")
	email := os.Getenv("JIT_AUTHOR_EMAIL")
	author := newAuthor(name, email)
	var message string
	fmt.Scanln(&message)

	ref := GitCommon.newRef(gitPath)
	parent := ref.readHead()

	commit := GitCommon.NewCommit(object.id, author, message, time.Now(), parent)
	commitObject := GitCommon.Object{kind: "commit", content: []byte(commit.toString())}
	database.write(commitObject)
	headPath := path2.Join(gitPath, "HEAD")
	file, _ := os.OpenFile(headPath, os.O_CREATE|os.O_WRONLY, 0777)
	file.Write([]byte("(root-commit) " + commitObject.id + commit.treeId))
}

func handleStatus() {

}

func main() {
	//commandDispatcher("commit")
	//writeToDisk()

}
