package packUtils

// unpacked

type Unpacked struct {
	entry Entry
	data  []byte
}

type Window struct {
	objects []Unpacked
	offset  int64
}

func NewWindow(size int64) *Window {
	objects := make([]Unpacked, size)
	return &Window{objects: objects, offset: 0}
}

func (window *Window) wrap(offset int64) int64 {
	return offset % int64(len(window.objects))
}

func (window *Window) each() []Unpacked {
	cursor := window.wrap(window.offset - 2)
	limit := window.wrap(window.offset - 1)
	result := []Unpacked{}
	for {
		if cursor == limit {
			break
		}
		unpacked := window.objects[cursor]
		result = append(result, unpacked)
		cursor = window.wrap(cursor - 1)
	}
	return result
}
func (window *Window) add(entry Entry, data []byte) *Unpacked {
	unpacked := Unpacked{entry: entry, data: data}
	window.objects[window.offset] = unpacked
	return &unpacked

}
