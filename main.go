package main

import (
	"fmt"
	"io/fs"
	"os"
	path2 "path"
)

//objects folder is used to
//refs folder is used to
func commandDispatcher(command string) {
	path, _ := os.Getwd()
	if command == "init"{
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
	if command == "commit" {
		gitPath := path2.Join(path,".git")
		objectsPath := path2.Join(gitPath, "objects")
		fmt.Println(objectsPath)
		workspace := Workspace{path: path,
		ignore: []string{".", "..", ".git"}}
		workspace.listFiles()
	}




}
func main(){
	commandDispatcher("commit")


}