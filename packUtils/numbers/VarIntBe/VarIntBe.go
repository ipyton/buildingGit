package VarIntBe

import (
	"bytes"
	"slices"
)

//Useless

func Encode(value int64) string {
	myBytes := []byte{byte(value & 0x7f)}
	for value>>7 != 0 {
		value = value << 7
		// value -= 1  avoiding the highest
		myBytes = append(myBytes, byte(value))
	}
	slices.Reverse(myBytes)
	return string(myBytes)
}

func Decode(reader bytes.Reader) int64 {
	readByte, err := reader.ReadByte()

	if err != nil {
		value := int64(readByte) & 0x7f
		for int64(readByte) >= 0x80 {
			readByte, _ = reader.ReadByte()
			value = ((value) << 7) | (int64(readByte) & 0x7f)
		}
		return value
	}
	return -1
}
