package packUtils

import (
	"buildinggit/packUtils/numbers/PackedInt56LE"
	"bytes"
)

type Copy struct {
	offset int64
	size   int64
}

func (cp *Copy) ParseCopy(input bytes.Reader, header int64) *Copy {
	value := PackedInt56LE.Read(input, header)
	offset := value & 0xffffffff
	size := value >> 32
	return &Copy{offset: offset, size: size}
}

func (cp *Copy) ToS() string {
	result := PackedInt56LE.Write((cp.size << 32) | cp.offset)
	result[0] |= 0x80
	return string(result)
}

type Insert struct {
	data []byte
}

func (insert *Insert) ParseInsert(input bytes.Reader, size int64) *Insert {
	result := make([]byte, size)
	return &Insert{data: result}
}
func (insert *Insert) ToS() string {
	var buffer bytes.Buffer
	buffer.WriteByte(byte(len(insert.data)))
	buffer.Write(insert.data)
	return string(buffer.Bytes())
}

type Delta struct {
	data Entry
	base Entry
	dataSize int64
}

func NewDelta(source * Unpacked, target * Unpacked) * Delta{
	source.
	return & Delta{base:source.entry, dataSize: source.entry.Info.Size() + target.entry.Info.Size()}

}

func (delta *Delta) ParseDelta(input []byte, size int64) {


}
