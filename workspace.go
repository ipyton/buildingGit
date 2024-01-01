package main

import (
	"os"
	"path"
)

type Workspace struct {
	path string
	ignore []string
}

func (workspace Workspace) init(path string, ignore []string)  {
	workspace.ignore = ignore
	workspace.path = path
}


func (workspace Workspace) listFiles() []string  {
	var result []string
	queue := make([]string, 0)
	queue = append(queue, workspace.path)
	for {
		if 0 == len(queue){
			break
		}
		basePath := queue[0]
		stat, _ := os.Stat(basePath)
		if stat.IsDir() {
			dirs, _ := os.ReadDir(basePath)
			for _,dir := range dirs {
				queue = append(queue, path.Join(basePath, dir.Name()))
			}
		} else {
			result = append(result, stat.Name())
		}
		queue = queue[1:]
	}


	return result
}

func readFile(path string) []byte{
	file, _ := os.ReadFile(path)
	return file

}

func statFile(path string) os.FileInfo{
	stat, _ := os.Stat(path)
	return stat
}
