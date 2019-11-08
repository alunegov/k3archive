package models

type FileDataMapper interface {
	GetFileInfo(name string) (uint8, string, error)
}
