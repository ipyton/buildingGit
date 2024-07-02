package database

type Blob struct {
	data string
}

func (b *Blob) ToS() string {
	return b.data
}

func (b *Blob) Parse(data string) {
	b.data = data
}

func (Blob) Type() string {
	return "blob"
}
