package chunks

import (
	"testing"

	"github.com/stretchr/testify/assert"

	_chunksModels "github.com/alunegov/k3archive/models/chunks"
)

func TestRosChunkFileMapperLoad(t *testing.T) {
	// TODO: pass ChunkMappers to parse ChunksData
	sut := NewRosChunkFileMapper(nil)

	res, err := sut.Load("../../testdata/Signal/7")
	if !assert.NoError(t, err) {
		return
	}

	if assert.NotNil(t, res) {
		ref := &_chunksModels.RosChunkFile{
			Hdr: _chunksModels.RosChunkFileHeader1{
				ChunkFileId: 0x1312,
				MaxChunks:   5,
				NumChunks:   5,
			},
			ChunksInfo: []_chunksModels.RosChunkFileHeader2{
				{_chunksModels.ChunkComment, 1, 71, 26},
				{_chunksModels.ChunkSignalSpec, 1, 97, 31},
				{_chunksModels.ChunkSignalHeader, 1, 128, 15},
				{_chunksModels.ChunkSignalHeader, 1, 143, 15},
				{_chunksModels.ChunkSignalData, 1, 158, 3200},
			},
			ChunksData: []_chunksModels.RosChunkFileData{
				nil,
				nil,
				nil,
				nil,
				nil,
			},
		}
		assert.Equal(t, ref, res)
	}
}
