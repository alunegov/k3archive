package chunks

const (
	ChunkSignalSpec      = 51
	ChunkSignalHeader    = 52
	ChunkSignalData      = 53
	ChunkSignalBalancing = 54
)

type SignalHeaderChunkV1 struct {
	RegChannelNum uint8
	DataType      uint8
	DataUnits     uint8
	PointsCount   uint32
	DX            float32
	Coeff         float32
}
