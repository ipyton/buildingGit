package Common

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

func (workspace *Workspace) init(path string) {
	workspace.ignore[".gitignore"] = true
	workspace.ignore["."] = true
	workspace.ignore[".."] = true
	workspace.path = path
}

func (workspace *Workspace) ListFiles(targetPath string) []string {
	//get all files (excluding directory)recursively in target path.
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

func (workspace *Workspace) listDirs(targetPath string) map[string]os.FileInfo {
	dir, _ := os.ReadDir(targetPath)
	hashmap := make(map[string]os.FileInfo)
	for _, file := range dir {
		if !workspace.ignore[file.Name()] {
			hashmap[file.Name()], _ = file.Info()
		}
	}
	return hashmap
}

func (workspace *Workspace) ReadFile(target string) []byte {
	file, _ := os.ReadFile(path.Join(workspace.path, target))
	return file
}

func (workspace *Workspace) writeFile(path string, data string, mode os.FileMode, mkdir bool) {

}

func (workspace *Workspace) removeDirectory(target string) {
	// remove all things including directory
	err := os.RemoveAll(target)
	if err != nil {
		println("error while removing")
	}
}

func (workspace *Workspace) Remove(target string) {
	// remove a single file
	targetPath := path.Join(workspace.path, target)
	os.Remove(targetPath)

	//If the folder contains the file, after deletion it became to empty, delete its parent folder.
	dirs, _ := os.ReadDir(targetPath)
	for _, dir := range dirs {
		if dir.IsDir() {
			workspace.removeDirectory(dir.Name())
		}
	}

}

func (workspace *Workspace) StatFile(targetPath string) os.FileInfo {
	stat, _ := os.Stat(path.Join(workspace.path, targetPath))
	return stat
}
