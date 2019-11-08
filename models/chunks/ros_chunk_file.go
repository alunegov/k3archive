package chunks

type RosChunkFile struct {
	Hdr        RosChunkFileHeader1
	ChunksInfo []RosChunkFileHeader2
	ChunksData []RosChunkFileData
}

type RosChunkFileHeader1 struct {
	ChunkFileId uint16
	MaxChunks   uint16
	NumChunks   uint16
}

type RosChunkFileHeader2 struct {
	ChunkId  uint32
	ChunkVer uint8
	Offset   uint32
	Size     uint32
}

type RosChunkFileData interface{}

func (it *RosChunkFile) GetChunksData(id int) []RosChunkFileData {
	res := make([]RosChunkFileData, 0, len(it.ChunksInfo))
	for i, info := range it.ChunksInfo {
		if int(info.ChunkId) == id {
			res = append(res, it.ChunksData[i])
		}
	}
	return res
}
