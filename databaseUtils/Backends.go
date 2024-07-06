package DatabaseUtils

import (
	"buildinggit/Base"
	"go/types"
)

type Backends struct {
	loose    Loose
	PathName string
	Stores   []Base.Store
}

func NewBackends(loose Loose, pathName string) *Backends {
	return &Backends{loose: Loose{pathname: pathName}, PathName: pathName}
}

func (backends *Backends) Reload() (Backends, error) {

}

func (backends *Backends) packPath() error {}

func (backends *Backends) LoadRaw(oid string) *Raw {
	for i := range backends.Stores {
		raw := backends.Stores[i].LoadRaw(oid)
		if raw != nil {
			return raw
		}
	}
	return nil

}

func (backends *Backends) LoadInfo(oid string) *types.Info {
	for i := range backends.Stores {
		raw := backends.Stores[i].LoadRaw(oid)
		if raw != nil {
			return raw
		}
	}
	return nil
}

func (backends *Backends) PrefixMatch() {

}

func (backends *Backends) SaveRaw(oid [20]byte) {

}

func (backends *Backends) Packed() {}
