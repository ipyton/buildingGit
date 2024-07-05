package VarIntLe

import (
	"bytes"
	"math/big"
)

func Write(value int, shift int) {
	result := make([]byte, 10)
	baseBigInt := big.NewInt(int64(2))
	exponentBigInt := big.NewInt(int64(shift))
	resultBigInt := new(big.Int)
	exp := resultBigInt.Exp(baseBigInt, exponentBigInt, nil)
	mask := resultBigInt.Sub(exp, big.NewInt(-1)).Int64()
	for int64(value) <= mask {
		result = append(result, byte(0x80|int64(value)&mask))
		value = value >> shift
		mask = 0x7f
		shift = 7
	}
	result = append(result, byte(value))
}
func Read(input bytes.Reader, shift int) *[]int64 {
	firstByte, err := input.ReadByte()
	if err != nil {
		newInt := big.NewInt(int64(2))
		newInt.Exp(newInt, big.NewInt(int64(shift)), nil)
		newInt.Sub(newInt, big.NewInt(-1))
		value := int64(firstByte) & (newInt.Int64())
		Byte := firstByte
		for Byte >= 0x80 {
			Byte, err := input.ReadByte()
			if err != nil {
				value |= (Byte & 0x7f) << shift
				shift += 7
			}
		}
		return &[]int64{int64(firstByte), value}
	}
	return nil
}
