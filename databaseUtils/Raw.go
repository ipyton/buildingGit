package DatabaseUtils

import (
	"io"
)

type Raw struct {
	Type string
	Size int
	Data io.Reader
}
