package main


type Blob struct {
	size int
	content []byte
	t string
	id string
}

func (Blob) newBlob(content []byte) Blob {
	return Blob{size:len(content),
	content: content,
	t:"object",id:""}
}
