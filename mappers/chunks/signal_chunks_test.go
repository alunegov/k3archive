package chunks

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	_chunksModels "github.com/alunegov/k3archive/models/chunks"
)

func TestSignalHeaderChunkMapperLoad(t *testing.T) {
	sut := NewSignalHeaderChunkMapper()

	f, err := os.Open("../../testdata/Signal/7")
	if !assert.NoError(t, err) {
		return
	}
	defer func(f_ *os.File) {
		_ = f_.Close()
	}(f)

	res := sut.Load(f, &_chunksModels.RosChunkFileHeader2{Offset: 143})

	if assert.NotNil(t, res) {
		switch o := res.(type) {
		case *_chunksModels.SignalHeaderChunkV1:
			ref := &_chunksModels.SignalHeaderChunkV1{
				RegChannelNum: 1,
				DataType:      _chunksModels.DataType_Spectrum,
				DataUnits:     _chunksModels.DataUnits_Marker,
				PointsCount:   400,
				DX:            2.5,
				Coeff:         1.0,
			}
			assert.Equal(t, ref, o)
		default:
			assert.Fail(t, "unsupp type")
		}
	}
}
