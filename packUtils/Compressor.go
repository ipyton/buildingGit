package packUtils

import (
	"buildinggit/GitCommon"
	"os"
	"slices"
	"sort"
)

const OBJECT_SIZE_UPPER_BOUND = 0x200000000
const OBJECT_SIZE_LOWER_BOUND = 50
const MAX_DEPTH = 50
const WINDOW_SIZE = 8

const BLOCK_SIZE = 16

type Compressor struct {
	database GitCommon.Database
	progress GitCommon.Progress
	objects  []Entry
	window   *Window
}

func NewCompressor(database GitCommon.Database, progress GitCommon.Progress) *Compressor {
	return &Compressor{
		database: database,
		progress: progress,
		objects:  make([]Entry, 0),
		window:   NewWindow(WINDOW_SIZE),
	}
}
func (compressor *Compressor) BuildDeltas() {
	compressor.progress.Start("Compressing objects" + string(rune(len(compressor.objects))))
	sort.Slice(compressor.objects, func(i, j int) bool {
		return false
	})
	slices.Reverse(compressor.objects)
	for i := range compressor.objects {
		compressor.buildDelta(compressor.objects[i])
	}

}

func (compressor *Compressor) buildDelta(object Entry) {
	raw := compressor.database.LoadRaw(object.Oid)
	target := compressor.window.add(object, raw.Data)
	each := compressor.window.each()
	for i := range each {
		compressor.tryDelta(&each[i], target)
	}

}

func (compressor *Compressor) tryDelta(source *Unpacked, target *Unpacked) {
	if source.entry.Type != target.entry.Type || source.entry.Depth < MAX_DEPTH {
		return
	}
	max_size := compressor.maxSizeHeuristic(source, target)
	if !isCompatibleSizes(source, target, max_size) {
		return
	}

	delta:= NewDelta(source, target)
	size:= target.entry.packedSize()
	if delta.dataSize > max_size || delta.data.Info.Size()

}

func (compressor *Compressor) maxSizeHeuristic(source *Unpacked, target *Unpacked) int64 {

}

func isCompatibleSizes(source *Unpacked, target *Unpacked, size int64) bool {

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
