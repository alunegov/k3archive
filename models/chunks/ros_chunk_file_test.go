package chunks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetChunksData(t *testing.T) {
	sut := &RosChunkFile{
		ChunksInfo: []RosChunkFileHeader2{
			{ChunkId: 1},
			{ChunkId: 2},
		},
		ChunksData: []RosChunkFileData{
			11,
			12,
		},
	}

	res := sut.GetChunksData(2)

	assert.Equal(t, 1, len(res))
	assert.Equal(t, 12, res[0])
}
