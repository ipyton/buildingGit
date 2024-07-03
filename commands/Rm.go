package commands

import (
	repository2 "buildinggit/repository"
	"flag"
)

const BOTH_CHANGED = "staged content different from both the file and the HEAD"
const INDEX_CHANGED = "changes staged in the index"
const WORKSPACE_CHANGED = "local modifications"

type Rm struct {
	Base
	options     map[string]bool
	headOid     [20]byte
	inspector   repository2.Inspector
	uncommitted []string
	unstaged    []string
	bothChanged []string
	paths       []string
}

func (rm *Rm) expandPath(path string) []string {
	//used to make a path to several directories
	if rm.Repo.Index.IsTrackedDirectory(path) {
		if rm.options["recursive"] {
			return rm.Repo.Index.ChildPaths(path)
		}
	}
	if rm.Repo.Index.IsTrackedFile(path) {
		return []string{path}
	}
	return []string{}
}

func (rm *Rm) defineOptions() {
	var cached = flag.Bool("--cached", false, "if the file is cached")
	var force = flag.Bool("-f", false, "force remove")
	var recursive = flag.Bool("-r", false, "recursively remove")
	flag.Parse()
	rm.options["cached"] = *cached
	rm.options["force"] = *force
	rm.options["recursive"] = *recursive
}

func (rm *Rm) run() {
	rm.Repo.Index.LoadForUpdate()
	rm.inspector = repository2.Inspector{Repo: rm.Repo}
	paths := []string{}
	for i := 0; i < len(rm.args); i++ {
		paths = append(paths, rm.expandPath(rm.args[i])...)
	}
	rm.args = paths
	for i := 0; i < len(rm.args); i++ {
		rm.planRemoval(rm.args[i])
	}

	for i := 0; i < len(rm.args); i++ {
		rm.remove(rm.args[i])
	}

	rm.Repo.Index.WriteUpdates()
}

func (rm *Rm) planRemoval(targetPath string) {
	if rm.options["force"] {
		return
	}
	stat := rm.Repo.Workspace.StatFile(targetPath)
	item := rm.Repo.Database.LoadTreeEntry(rm.headOid, targetPath)
	entry := rm.Repo.Index.EntryForPath()

	rm.inspector.CompareTreeToIndex(item, stat)
	rm.inspector.CompareIndexToWorkspace()

	rm.Repo.Index.Remove(targetPath)

}

func (rm *Rm) remove(target string) {
	rm.Repo.Index.Remove(target)
	if !rm.options["cached"] {
		rm.Repo.Workspace.remove(target)
	}

}
