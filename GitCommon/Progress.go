package GitCommon

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"syscall"
	"time"
)

var UNITS = []string{"B", "KiB", "MiB", "GiB"}
var SCALE = 1024.0

type Progress struct {
	out     *os.File
	message string
	count   int64
	size    int64
	total   int64
	writeAt time.Time
}

func isTerminal(fd uintptr) bool {
	var mode uint32
	// Get the handle for the file descriptor
	handle := syscall.Handle(fd)
	// Try to get the console mode for the handle
	err := syscall.GetConsoleMode(handle, &mode)
	return err == nil
}

func (progress *Progress) Start(message string) {
	os.Stdout.Fd()
	if os.Getenv("NO_PROGRESS") == "" || isTerminal(os.Stdout.Fd()) {
		return
	}
	progress.message = message

}

func (progress *Progress) Tick(size int64) {
	if progress.message == "" {
		return
	}
	progress.count++
	progress.size = size
	if time.Now().Before(progress.writeAt.Add(0.05)) {
		return
	}
	progress.writeAt = time.Now()
	progress.ClearLine()
	progress.out.Write([]byte(progress.StatusLine()))
}

func (progress *Progress) Stop() {
	if progress.message == "" {
		return
	}
	progress.total = progress.count
	progress.ClearLine()
	progress.out.Write([]byte(progress.StatusLine()))
	progress.message = ""
}

func (progress *Progress) ClearLine() {
	progress.out.Write([]byte("\033[2K\033[0G"))
}

func (progress *Progress) StatusLine() string {
	line := fmt.Sprintf("%s %d", progress.message, progress.FormatCount())
	return line
}

func (progress *Progress) FormatCount() string {
	percent := 100.0
	if progress.total > 0 {
		percent = float64(100 * progress.count / progress.total)
		return fmt.Sprintf("%.2f %% %d /%d", percent, progress.total, progress.count)
	}
	return strconv.Itoa(int(progress.count))
}

func (progress *Progress) FormatSize(size int64) string {
	power := int(math.Floor(math.Log(float64(progress.size)) / math.Log(1024)))
	scaled := float64(progress.size) / math.Pow(SCALE, float64(power))
	return fmt.Sprintf("%.2f %s", scaled, UNITS[power])
}
