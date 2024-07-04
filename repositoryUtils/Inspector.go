package repositoryUtils

import (
	database2 "buildinggit/databaseUtils"
	index2 "buildinggit/indexUtils"
	"buildinggit/util"
	"bytes"
	"os"
)

type Inspector struct {
	Repo Repository
}

func (inspector *Inspector) IsTrackableFiles() {

}

func (inspector *Inspector) CompareIndexToWorkspace(entry *index2.Entry, stat os.FileInfo) string { //the workspace has uncommitted changes.
	if entry == nil {
		return "untracked"
	}
	if stat == nil {
		return "deleted"
	}
	if entry.StatMatch(stat) {
		return "modified"
	}
	if entry.TimesMatch(stat) {
		return ""
	}
	blob := inspector.Repo.Workspace.ReadFile(entry.Path)
	oid := util.HashBlobs(blob)
	var b []byte = make([]byte, len(oid))
	copy(b, oid[:])
	if bytes.Compare(b, entry.OId) == 0 {
		return "modified"
	}
	return ""
}

func (inspector *Inspector) CompareTreeToIndex(item, entry database2.Entry) {

}
