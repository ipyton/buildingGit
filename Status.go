package main

type Status struct {
	untracked []string
	changes map[string] []string
	Base
}


func LongStatus() map[string]string {
	m := map[string]string {
		"added" : "new file:",
		"deleted" : "deleted:",
		"modified" : "modified:",
	}
	return m
}

func ShortStatus() map[string]string {
	m := map[string]string {
		"added" : "A",
		"deleted" : "D",
		"modified" : "M",
	}
	return m
}


func (status Status) run() {
	status.repository.index.loadForUpdate()

	status.scanWorkspace()
	status.loadHeadTree()

	//status.detectWorkspaceChanges()
	status.checkIndexEntries()

	status.repository.index.writeUpdates()
}

func (status Status) printResults() {

}

func (status Status) loadHeadTree() {
	headOid :=status.repository.refs.readHead()
	if len(headOid) == 0 {
		return
	}

}

func (status Status) printLongFormat() {

}

func (status Status) printChanges() {

}

func (status Status) printCommitStatus() {

}

func (status Status) printPorcelainFormat() {

}

func (status Status) statusFor(path string) string {
	changes := status.changes[path]
	left := ""
	right := ""
	for _, value := range changes {
		if value == "index_add" {
			left = "A"
		}
		if value == "workspace_deleted" {
			 right = "D"
		} else if value == "workspace_modified" {
			right = "M"
		}
	}
	return left + right
}

func (status Status) scanWorkspace() {

}

func (status Status) detectWorkspaceChanges() {

}

func (status Status) checkIndexEntries(entry Entry) {

}

func (status Status) recordChange() {

}
