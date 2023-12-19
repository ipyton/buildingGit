package main

import (
	"fmt"
	"os"
)

type Workspace struct {
	path string
	ignore []string
}

func (receiver Workspace) init(path string, ignore []string)  {
	receiver.ignore = ignore
	receiver.path = path
}

func (receiver Workspace) listFiles()  {
	dirs, _ := os.ReadDir(receiver.path)
	for _, dir :=range dirs{
		fmt.Println(dir)

	}
}