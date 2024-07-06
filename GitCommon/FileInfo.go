package GitCommon

import (
	"strconv"
	"strings"
	"time"
)

type FileInfo struct {
	Name     string
	Size     int64
	Mode     uint32
	ModeTime time.Time
	IsDir    bool
}

func NewFileInfo(name string, size int64, mode uint32, modeTime time.Time, isDir bool) *FileInfo {
	return &FileInfo{Name: name, Size: size, Mode: mode, ModeTime: modeTime, IsDir: isDir}
}

func ParseFileInfo(info string) *FileInfo {
	split := strings.Split(info, "_")
	size, err := strconv.ParseInt(split[1], 0, 64)
	if err == nil {
		return nil
	}
	parseUint, err := strconv.ParseUint(split[2], 0, 32)

	datetime, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", split[3])
	isDir, err := strconv.ParseBool(split[4])

	return NewFileInfo(split[0], size, uint32(parseUint), datetime, isDir)
}
