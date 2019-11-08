package mappers

import (
	_chunksMappers "github.com/alunegov/k3archive/mappers/chunks"
	_chunksModels "github.com/alunegov/k3archive/models/chunks"
)

const rmsBaseOpts = _chunksModels.OptsTypeSkz << _chunksModels.OptsTypeShift

// implements models.FileDataMapper
type RmsMapper struct {
	mapper *_chunksMappers.RosChunkFileMapper
}

func NewRmsMapper(mapper *_chunksMappers.RosChunkFileMapper) *RmsMapper {
	return &RmsMapper{
		mapper: mapper,
	}
}

func (it *RmsMapper) GetFileInfo(name string) (uint8, string, error) {
	file, err := it.mapper.Load(name)
	if err != nil {
		return 0, "", err
	}

	comment, err := getComment(file)
	if err != nil {
		return 0, "", err
	}

	return rmsBaseOpts, comment, nil
}
