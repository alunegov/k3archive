package chunks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRosChunkFileMapperLoad(t *testing.T) {
	sut := NewRosChunkFileMapper(nil)

	res, err := sut.Load("../../testdata/Signal/7")
	if !assert.NoError(t, err) {
		return
	}

	if assert.NotNil(t, res) {
		assert.Equal(t, 5, len(res.ChunksInfo))
		assert.Equal(t, 5, len(res.ChunksData))
	}
}
