package models

import (
	"os"
	"time"
)

type FileInfo struct {
	Name    string
	Size    int64
	ModTime time.Time
	Path    string
	Opts    uint8
	Comment string
}

func NewFileInfo(sysFileInfo os.FileInfo, path string) *FileInfo {
	return &FileInfo{
		Name:    sysFileInfo.Name(),
		Size:    sysFileInfo.Size(),
		ModTime: sysFileInfo.ModTime(),
		Path:    path + "/" + sysFileInfo.Name(),
		Opts:    0,
		Comment: "",
	}
}
