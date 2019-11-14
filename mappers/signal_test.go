package mappers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/text/encoding/htmlindex"

	_chunksMappers "github.com/alunegov/k3archive/mappers/chunks"
	_chunksModels "github.com/alunegov/k3archive/models/chunks"
)

func TestSignalGetFileInfo(t *testing.T) {
	cases := []struct {
		fileName   string
		expError   bool
		expOpts    uint8
		expComment string
	}{
		{"../testdata/Signal/7", false, 0x14, "qweqwe"},
		{"../testdata/Signal/0", true, 0, ""}, // no file
	}

	// TODO: refactor _chunksMappers.RosChunkFileMapper to interface
	enc, err := htmlindex.Get("windows-1251")
	if !assert.NoError(t, err) {
		return
	}
	chunkMappers := _chunksMappers.ChunkMappers{
		_chunksModels.ChunkComment:      _chunksMappers.NewCommentChunkMapper(enc.NewDecoder()),
		_chunksModels.ChunkSignalHeader: _chunksMappers.NewSignalHeaderChunkMapper(),
	}
	rosChunkFileMapper := _chunksMappers.NewRosChunkFileMapper(chunkMappers)

	sut := NewSignalMapper(rosChunkFileMapper)

	for _, c := range cases {
		opts, comment, err := sut.GetFileInfo(c.fileName)

		assert.Equal(t, c.expError, err != nil)
		assert.Equal(t, c.expOpts, opts)
		assert.Equal(t, c.expComment, comment)
	}
}
