package chunks

import (
	"encoding/binary"
	"os"

	_chunksModels "github.com/alunegov/k3archive/models/chunks"
)

// implements ChunkMapper
type SignalHeaderChunkMapper struct{}

func NewSignalHeaderChunkMapper() *SignalHeaderChunkMapper {
	return &SignalHeaderChunkMapper{}
}

func (it *SignalHeaderChunkMapper) Load(file *os.File, chunkInfo *_chunksModels.RosChunkFileHeader2) _chunksModels.RosChunkFileData {
	_, _ = file.Seek(int64(chunkInfo.Offset), 0)

	// V1
	res := &_chunksModels.SignalHeaderChunkV1{}
	_ = binary.Read(file, binary.LittleEndian, res)

	return res
}
