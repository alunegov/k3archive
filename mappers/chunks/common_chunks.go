package chunks

import (
	"bytes"
	"os"

	"golang.org/x/text/encoding"

	_chunksModels "github.com/alunegov/k3archive/models/chunks"
)

// implements ChunkMapper
type CommentChunkMapper struct {
	textDec *encoding.Decoder
}

func NewCommentChunkMapper(textDec *encoding.Decoder) *CommentChunkMapper {
	return &CommentChunkMapper{
		textDec: textDec,
	}
}

func (it *CommentChunkMapper) Load(file *os.File, chunkInfo *_chunksModels.RosChunkFileHeader2) _chunksModels.RosChunkFileData {
	// V1
	b := make([]byte, 26)
	_, _ = file.ReadAt(b, int64(chunkInfo.Offset))

	n := bytes.IndexByte(b, 0)
	if n == -1 {
		n = len(b)
	}

	b2, _ := it.textDec.Bytes(b[:n])

	return &_chunksModels.CommentChunkV1{
		Text: string(b2),
	}
}
