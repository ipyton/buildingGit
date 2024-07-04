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

func (inspector *Inspector) CompareIndexToWorkspace(item *index2.Entry, entry *database2.Entry) string { //the workspace has uncommitted changes.
	///item is for staged, entry is for committed
	if item == nil && entry == nil {
		return ""
	}
	if item == nil {
		return "added"
	}
	if entry == nil {
		return "deleted"
	}

	if item.Mode != entry.Mode || bytes.Compare(item.Oid, entry.Oid) == 0 {
		return "modified"
	}
	return ""
}

func (inspector *Inspector) CompareTreeToIndex(entry *index2.Entry, stat os.FileInfo) string {
	// stat is for workspace, entry is for staged changes
	// stat is for workspace, entry is for staged changes
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
	if bytes.Compare(b, entry.Oid) == 0 {
		return "modified"
	}
	return ""

}
