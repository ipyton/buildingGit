package main

import (
	"strconv"
	"strings"
	"time"
)

type FileInfo struct {
	name string
	size int64
	mode uint32
	modeTime time.Time
	isDir bool
}


func NewFileInfo(name string, size int64, mode uint32, modeTime time.Time, isDir bool) *FileInfo {
	return & FileInfo{name: name,size: size,mode: mode, modeTime: modeTime, isDir: isDir}
}


func parseFileInfo(info string) * FileInfo {
	split := strings.Split(info, "_")
	size, err := strconv.ParseInt(split[1],0,64)
	if err == nil {
		return nil
	}
	parseUint, err := strconv.ParseUint(split[2], 0, 32)

	datetime, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", split[3])
	isDir, err := strconv.ParseBool(split[4])

	return NewFileInfo(split[0], size, uint32(parseUint), datetime, isDir)
}