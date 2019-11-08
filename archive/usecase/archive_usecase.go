package usecase

import (
	"os"

	"github.com/alunegov/k3archive/models"
)

// implements archive.Usecase
type ArchiveUsecase struct {
	fsRoot string
	path   string
	mapper models.FileDataMapper
}

func NewArchiveUsecase(fsRoot string, path string, mapper models.FileDataMapper) *ArchiveUsecase {
	return &ArchiveUsecase{
		fsRoot: fsRoot,
		path:   path,
		mapper: mapper,
	}
}

func (it *ArchiveUsecase) GetDir() ([]*models.FileInfo, error) {
	sysFile, err := os.Open(it.fsRoot + "/" + it.path)
	if err != nil {
		return nil, err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(sysFile)

	// TODO: Readdir could return files info AND error
	sysFilesInfo, err := sysFile.Readdir(-1)
	if err != nil {
		return nil, err
	}

	res := make([]*models.FileInfo, 0)

	for _, sysFileInfo := range sysFilesInfo {
		if sysFileInfo.IsDir() {
			continue
		}

		res = append(res, models.NewFileInfo(sysFileInfo, "/"+it.path))
	}

	for _, r := range res {
		r.Opts, r.Comment, err = it.mapper.GetFileInfo(it.fsRoot + r.Path)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

func (it *ArchiveUsecase) GetFile(id string) (*models.FileInfo, error) {
	sysFileInfo, err := os.Stat(it.fsRoot + "/" + it.path + "/" + id)
	if err != nil {
		return nil, err
	}

	res := models.NewFileInfo(sysFileInfo, "/"+it.path)

	res.Opts, res.Comment, err = it.mapper.GetFileInfo(it.fsRoot + res.Path)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (it *ArchiveUsecase) DeleteFile(id string) error {
	return os.Remove(it.fsRoot + "/" + it.path + "/" + id)
}
