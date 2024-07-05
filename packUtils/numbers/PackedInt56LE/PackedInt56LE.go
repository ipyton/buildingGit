package PackedInt56LE

import (
	"bytes"
	"os"
)

func Write(value int64) []byte {
	result := make([]byte, 8)
	result = append(result, 0)
	for i := 0; i <= 7; i++ {
		tmpByte := value >> (8 * i) & 0xff
		if tmpByte == 0 {
			continue
		}
		result[0] |= 1 << i
		result = append(result, tmpByte)
	}
	return result
}

func Read(reader bytes.Reader, header int64) int64 {
	var value int64 = 0
	for i := 0; i <= 7; i++ {
		if header&(1<<i) != 0 {
			readByte, err := reader.ReadByte()
			if err != nil {
				value |= int64(readByte) << (8 * i)
			} else {
				println(err)
				os.Exit(-1)
			}
		}
	}
	return value
}
