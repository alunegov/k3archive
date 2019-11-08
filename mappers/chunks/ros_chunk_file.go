package chunks

import (
	"encoding/binary"
	"fmt"
	"os"

	_chunksModels "github.com/alunegov/k3archive/models/chunks"
)

type RosChunkFileMapper struct {
	mappers ChunkMappers
}

type ChunkMappers map[uint32]ChunkMapper

type ChunkMapper interface {
	Load(file *os.File, chunkInfo *_chunksModels.RosChunkFileHeader2) _chunksModels.RosChunkFileData
}

func NewRosChunkFileMapper(mappers ChunkMappers) *RosChunkFileMapper {
	return &RosChunkFileMapper{
		mappers: mappers,
	}
}

func (it *RosChunkFileMapper) Load(name string) (*_chunksModels.RosChunkFile, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(file)

	hdr := &_chunksModels.RosChunkFileHeader1{}
	binary.Read(file, binary.LittleEndian, hdr)

	chunksInfo := make([]_chunksModels.RosChunkFileHeader2, hdr.MaxChunks)
	binary.Read(file, binary.LittleEndian, chunksInfo)

	chunksData := make([]_chunksModels.RosChunkFileData, hdr.NumChunks)
	for i := range chunksData {
		mapper, ok := it.mappers[chunksInfo[i].ChunkId]
		if !ok {
			continue
		}
		chunksData[i] = mapper.Load(file, &chunksInfo[i])
	}

	res := &_chunksModels.RosChunkFile{
		Hdr:        *hdr,
		ChunksInfo: chunksInfo,
		ChunksData: chunksData,
	}

	fmt.Printf("%v\n", res)
	for i := range res.ChunksData {
		fmt.Printf("%v\n", res.ChunksData[i])
	}

	return res, nil
}
