package commands

const LOCKED_INDEX_ERROR_MSG = "Another jit process seems to be running in this repository.\n      Please make sure all processes are terminated then try again.\n      If it still fails, a jit process may have crashed in this\n      repository earlier: remove the file manually to continue.\n"

type Add struct {
	Base
}

func (Add) run() {

}

func (Add) expandedPaths() {

}

func (Add) addToIndex(path string) {

}

func (Add) handleLockedIndex() {

}

func (Add) handleMissingFile() {

}

func (Add) handleUnreadableFile() {

}
