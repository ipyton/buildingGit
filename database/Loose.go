package database

import "path"

// a loose file stands for not compressed files in the project.
// packed file is used to compress files

type Loose struct {
	pathname string
}

func (loose *Loose) Has(oid [20]byte) {

}

func (loose *Loose) LoadInfo(oid [20]byte) {

}

func (loose *Loose) LoadRaw(oid [20]byte) {

}

func (loose *Loose) PrefixMatch(name string) {

}

func (loose *Loose) WriteObject(oid []byte, content string) {
	return
}

func (loose *Loose) ObjectPath(oid [20]byte) string {
	return path.Join(string(oid[0:2]), string(oid[2:]))
}

func (loose *Loose) ReadObjectHeader(oid [20]byte) {

}
