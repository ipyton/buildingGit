package databaseUtils

type Blob struct {
	Data string
	Oid  string
}

func (b *Blob) ToS() string {
	return b.Data
}

func (b *Blob) Parse(data string) {
	b.Data = data
}

func (Blob) Type() string {
	return "blob"
}
