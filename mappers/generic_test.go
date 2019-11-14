package mappers

import (
	"testing"

	"github.com/stretchr/testify/assert"

	_chunksModels "github.com/alunegov/k3archive/models/chunks"
)

func TestGetComment(t *testing.T) {
	cases := []struct {
		chunkId    uint32
		data       interface{}
		expError   bool
		expComment string
	}{
		{_chunksModels.ChunkComment, &_chunksModels.CommentChunkV1{"test тест"}, false, "test тест"},
		{_chunksModels.ChunkComment, nil, false, ""}, // got chunk but no data
		{_chunksModels.ChunkComment, 1, true, ""},    // got chunk but unsupp data
		{_chunksModels.ChunkRmsData, nil, false, ""}, // no chunk
	}

	for _, c := range cases {
		f := &_chunksModels.RosChunkFile{
			ChunksInfo: []_chunksModels.RosChunkFileHeader2{
				{ChunkId: c.chunkId},
			},
			ChunksData: []_chunksModels.RosChunkFileData{
				c.data,
			},
		}

		res, err := getComment(f)
		assert.Equal(t, c.expError, err != nil)
		assert.Equal(t, c.expComment, res)
	}
}
