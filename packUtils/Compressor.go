package packUtils

import (
	"os"
)

const BLOCK_SIZE = 16

type Compressor struct {
	database Database
}

func NewCompressor(database Database) *Compressor {

}

func (compresor *Compressor) CreateIndex(file os.File) {
	stat, err := file.Stat()
	index := make(map[string][]int)

	if err != nil {
		blocks := stat.Size() / BLOCK_SIZE
		for index := int64(0); index < blocks; index++ {
			offset := index * BLOCK_SIZE

		}
	}

}
