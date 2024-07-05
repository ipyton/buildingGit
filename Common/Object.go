package Common

type Object struct {
	size    int
	content []byte
	kind    string
	id      string
}

func newObject(content []byte, kind string) Object {
	return Object{size: len(content),
		content: content,
		kind:    kind,
		id:      ""}
}
