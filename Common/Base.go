package Common

import repository2 "buildinggit/repositoryUtils"

type Base struct {
	dir          string
	args         map[string]string
	repository   repository2.Repository
	parseOptions []string
}

func newBase(dir string, args []string) {

}
