package main

import (
	"os"
)

type Workspace struct {
	path string
	ignore []string
}

func (receiver Workspace) init(path string, ignore []string)  {
	receiver.ignore = ignore
	receiver.path = path
}

func (receiver Workspace) listFiles() []string  {
	dirs, _ := os.ReadDir(receiver.path)
	var ignored map[string]string
	for _, value := range receiver.ignore {
		ignored[value] = ""
	}

	var result []string
	for _, dir :=range dirs{
		if !dir.IsDir(){
			path:=dir.Name()
			if ignored[path] != ""{
				result = append(result, path)
			}
		}
	}
	return result
}

func readFile(path string) []byte{
	file, _ := os.ReadFile(path)
	return file

}