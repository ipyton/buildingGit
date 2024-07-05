package commands

import "buildinggit/databaseUtils"

const LOCKED_INDEX_ERROR_MSG = "Another jit process seems to be running in this repositoryUtils.\n      Please make sure all processes are terminated then try again.\n      If it still fails, a jit process may have crashed in this\n      repositoryUtils earlier: remove the file manually to continue.\n"

type Add struct {
	Base
}

func (add *Add) run() {
	add.Repo.Index.LoadForUpdate()
	paths := add.expandedPaths()
	for i := range paths {
		add.addToIndex(paths[i])
	}
	add.Repo.Index.WriteUpdates()
}

func (add *Add) expandedPaths() []string {
	result := make([]string, 10)
	for i := range add.args {
		result = append(result, add.Repo.Workspace.ListFiles(add.expandedPathname(add.args[i]))...)
	}
	return result
}

func (add *Add) addToIndex(path string) {
	file := add.Repo.Workspace.ReadFile(path)
	stat := add.Repo.Workspace.StatFile(path)
	blob := databaseUtils.Blob{Data: string(file)}
	add.Repo.Database.Store(blob) // store will give it a oid
	add.Repo.Index.Add(path, blob.Oid, stat)

}

func (add *Add) handleLockedIndex() {

}

func (add *Add) handleMissingFile() {

}

func (add *Add) handleUnreadableFile() {

}
