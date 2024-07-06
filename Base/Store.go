package Base

import (
	"buildinggit/DatabaseUtils"
	"go/types"
)

type Store interface {
	LoadRaw(string) *DatabaseUtils.Raw
	LoadInfo(string) *types.Info
}
