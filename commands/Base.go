package commands

import (
	"buildinggit/repositoryUtils"
	"os"
)

type Base struct {
	dir    string
	env    map[string]string
	args   []string
	status int
	stdout *os.File
	stderr *os.File
	stdin  *os.File
	Repo   repositoryUtils.Repository
}

func (base *Base) newBase(dir string, env map[string]string, args []string) Base {
	return Base{dir: dir, env: env, args: args}
}

func (base *Base) execute() {

}

func (base *Base) repo() {

}

func (base *Base) expandedPathname() {

}

func (base *Base) parseOptions() {

}

func (base *Base) defineOptions() {

}

func (base *Base) setupPager() {

}

func (base *Base) editFile() {

}

func (base *Base) editCommand() {

}

func (base *Base) fmt() {

}
func (base *Base) puts() {

}

func (base *Base) exit() {

}
