package DatabaseUtils

type Backends struct {
	loose    Loose
	PathName string
	Stores
}

func NewBackends(loose Loose, pathName string) *Backends {
	return &Backends{loose: Loose{pathname: pathName}, PathName: pathName}
}

func (backends *Backends) Reload() (Backends, error) {

}

func (backends *Backends) packPath() error {}

func (backends *Backends) LoadRaw(oid [20]byte) {

}

func (backends *Backends) LoadInfo() {

}

func (backends *Backends) PrefixMatch() {

}

func (backends *Backends) SaveRaw(oid [20]byte) {

}

func (backends *Backends) Packed() {}
