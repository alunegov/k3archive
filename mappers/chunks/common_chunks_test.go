package chunks

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/text/encoding/htmlindex"

	_chunksModels "github.com/alunegov/k3archive/models/chunks"
)

func TestCommentChunkMapperLoad(t *testing.T) {
	enc, err := htmlindex.Get("windows-1251")
	if !assert.NoError(t, err) {
		return
	}

	sut := NewCommentChunkMapper(enc.NewDecoder())

	f, err := os.Open("../../testdata/Rms/3")
	if !assert.NoError(t, err) {
		return
	}
	defer func(f_ *os.File) {
		_ = f_.Close()
	}(f)

	res := sut.Load(f, &_chunksModels.RosChunkFileHeader2{Offset: 45})

	if assert.NotNil(t, res) {
		switch o := res.(type) {
		case *_chunksModels.CommentChunkV1:
			assert.Equal(t, "dsfg укеук", o.Text)
		default:
			assert.Fail(t, "unsupp type")
		}
	}
}
