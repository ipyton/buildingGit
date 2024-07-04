package repositoryUtils

// this file is used to get all informations that git status command needed.
type Status struct {
	repository       Repository
	IsExecutable     bool
	inspector        Inspector
	changed          map[string]bool
	IndexChanges     map[string]string
	Conflicts        map[string]string
	WorkspaceChanges map[string]string
	UntrackedFiles   map[string]bool
}

func NewStatus(repository Repository, commitOid []byte) *Status {
	status := Status{repository: repository}
	status.scanWorkspace()
	status.checkIndexEntries()
	status.collectDeletedHeadFiles()
	return &status
}

func (status *Status) scanWorkspace() {

}

func (status *Status) checkIndexEntries() {

}

func (status *Status) checkConflicts() {

}

func (status *Status) collectDeletedHeadFiles() {

}
