package main

import (
	"os"
	"path"
)

type Workspace struct {
	path   string
	ignore map[string]bool
}

// manipulating the files in directory
//

func (workspace Workspace) init(path string) {
	workspace.ignore[".gitignore"] = true
	workspace.ignore["."] = true
	workspace.ignore[".."] = true
	workspace.path = path
}

func listFilesByPath(targetPath string) []string {
	//get all files recursively in target path.
	var result []string
	queue := make([]string, 0)
	queue = append(queue, targetPath)
	for {
		if 0 == len(queue) {
			break
		}
		basePath := queue[0]
		stat, _ := os.Stat(basePath)
		if stat.IsDir() {
			dirs, _ := os.ReadDir(basePath)
			for _, dir := range dirs {
				queue = append(queue, path.Join(basePath, dir.Name()))
			}
		} else {
			result = append(result, stat.Name())
		}
		queue = queue[1:]
	}
	return result
}

func (workspace Workspace) listFiles() []string {
	return listFilesByPath(workspace.path)
}

func (workspace Workspace) listDirs(targetPath string) map[string]os.FileInfo {
	dir, _ := os.ReadDir(targetPath)
	hashmap := make(map[string]os.FileInfo)
	for _, file := range dir {
		if !workspace.ignore[file.Name()] {
			hashmap[file.Name()], _ = file.Info()
		}
	}
	return hashmap
}

func (workspace Workspace) readFile(target string) []byte {
	file, _ := os.ReadFile(path.Join(workspace.path, target))
	return file
}

func (workspace Workspace) writeFile(path string, data string, mode os.FileMode, mkdir bool) {

}

func (workspace Workspace) removeFiles(target string) {
	err := os.RemoveAll(target)
	if err != nil {
		println("error while removing")
	}
}

func statFile(path string) os.FileInfo {
	stat, _ := os.Stat(path)
	return stat
}
