package repository

import "slices"

const messageForStaleFile = "Your local changes to the following files would be overwritten by checkout:\n",
const messageForStaleFilePost = "Please commit your changes or stash them before you switch branches.\n"
const staleDirectory = "Updating the following directories would lose untracked files in them:\n"
const untracked_overwritten = "The following untracked working tree files would be overwritten by checkout:"
const untracked_overwritten_post = "Please move or remove them before you switch branches."
const untracked_removed = "The following untracked working tree files would be removed by checkout:"
const untracked_removed_post = "Please move or remove them before you switch branches."

type Migration struct {
	repository Repository
	inspector Inspector
	create []string
	update []string
	delete []string
	mkdirs map[string]bool
	rmdirs map[string]bool
	diff.
}

func (migration *Migration) Init(repository Repository, inspector Inspector) {


}
func checkForConflict(path string, oldItem, newItem string) bool {

}


func (migration *Migration) applyChanges() {
	migration.planChanges()
	migration.updateWorkspace()
	migration.updateIndex()
}

func (migration * Migration) planChanges() {

}

func (migration * Migration) updateWorkspace() {

}

func (migration * Migration) updateIndex() {

}

func (migration *Migration) getBlobData(oid string ) {
	migration.repository.database.load(oid)
}
