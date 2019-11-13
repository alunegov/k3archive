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

	res := sut.Load(f, &_chunksModels.RosChunkFileHeader2{Offset: 0x8f})

	if assert.NotNil(t, res) {
		switch o := res.(type) {
		case *_chunksModels.SignalHeaderChunkV1:
			assert.Equal(t, uint8(0x1), o.RegChannelNum)
			assert.Equal(t, uint8(0x1), o.DataType)
			assert.Equal(t, uint8(0x3), o.DataUnits)
			assert.Equal(t, uint32(0x190), o.PointsCount)
			assert.Equal(t, float32(2.5), o.DX)
			assert.Equal(t, float32(1), o.Coeff)
		default:
			assert.Fail(t, "unsupp type")
		}
	}
}
