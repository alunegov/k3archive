package mappers

import (
	_chunksMappers "github.com/alunegov/k3archive/mappers/chunks"
	_chunksModels "github.com/alunegov/k3archive/models/chunks"
)

const signalBaseOpts uint8 = _chunksModels.OptsTypeSignal << _chunksModels.OptsTypeShift

// implements models.FileDataMapper
type SignalMapper struct {
	mapper *_chunksMappers.RosChunkFileMapper
}

func NewSignalMapper(mapper *_chunksMappers.RosChunkFileMapper) *SignalMapper {
	return &SignalMapper{
		mapper: mapper,
	}
}

func (it *SignalMapper) GetFileInfo(name string) (uint8, string, error) {
	file, err := it.mapper.Load(name)
	if err != nil {
		return 0, "", err
	}

	opts, err := it.getOpts(file)
	if err != nil {
		return 0, "", err
	}

	comment, err := getComment(file)
	if err != nil {
		return 0, "", err
	}

	return opts, comment, nil
}

func (it *SignalMapper) getOpts(file *_chunksModels.RosChunkFile) (uint8, error) {
	headers := file.GetChunksData(_chunksModels.ChunkSignalHeader)
	if len(headers) < 1 || headers[0] == nil {
		return signalBaseOpts, nil
	}

	// по первой шапке (как в К2)
	hdrOpts := uint8(0)
	switch hdr := headers[0].(type) {
	case *_chunksModels.SignalHeaderChunkV1:
		switch hdr.DataType {
		case _chunksModels.DataType_Signal:
			hdrOpts |= _chunksModels.OptsSignalTypeSignal << _chunksModels.OptsSignalTypeShift
		case _chunksModels.DataType_Spectrum:
			hdrOpts |= _chunksModels.OptsSignalTypeSpectr << _chunksModels.OptsSignalTypeShift
		}

		switch hdr.DataUnits {
		case _chunksModels.DataUnits_VibroAcceleration:
			hdrOpts |= _chunksModels.OptsSignalEdIzmAcc << _chunksModels.OptsSignalEdIzmShift
		case _chunksModels.DataUnits_VibroVelocity:
			hdrOpts |= _chunksModels.OptsSignalEdIzmVel << _chunksModels.OptsSignalEdIzmShift
		case _chunksModels.DataUnits_VibroDisplacement:
			hdrOpts |= _chunksModels.OptsSignalEdIzmDis << _chunksModels.OptsSignalEdIzmShift
		}
	}

	return signalBaseOpts | hdrOpts, nil
}
