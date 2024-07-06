package GitCommon

import (
	index2 "buildinggit/indexUtils"
)

type Status struct {
	untracked []string
	changes   map[string][]string
	main.Base
}

func LongStatus() map[string]string {
	m := map[string]string{
		"added":    "new file:",
		"deleted":  "deleted:",
		"modified": "modified:",
	}
	return m
}

func ShortStatus() map[string]string {
	m := map[string]string{
		"added":    "A",
		"deleted":  "D",
		"modified": "M",
	}
	return m
}

func check_index_entry(entry index2.Entry) {

}

func recordChange() {

}

func (status Status) run() {
	status.repository.Index.loadForUpdate()
	status.printResults()

}

func (status Status) loadHeadTree() {
	headOid := status.repository.refs.readHead()
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

func (status Status) printResults() {
	if status.args["format"] == "long" {
		status.printLongFormat()
	} else if status.args["format"] == "porcelain" {
		status.printPorcelainFormat()
	}
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
