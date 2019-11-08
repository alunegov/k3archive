package mappers

import (
	"errors"

	_chunksModels "github.com/alunegov/k3archive/models/chunks"
)

func getComment(file *_chunksModels.RosChunkFile) (string, error) {
	comments := file.GetChunksData(_chunksModels.ChunkComment)
	if len(comments) < 1 || comments[0] == nil {
		return "", nil // no comment
	}

	comment, ok := comments[0].(*_chunksModels.CommentChunkV1)
	if !ok {
		return "", errors.New("not *_chunksModels.CommentChunkV1")
	}

	return comment.Text, nil
}
