package chunks

const (
	ChunkRmsSpec = 101
	ChunkRmsData = 102
)

type RmsSpecChunkV1 struct {
	CalcedParamsStart     uint8
	CalcedParamsNext      []uint8
	ParamForVibroChannels []uint8
	ValuesPresence        []uint16
}
