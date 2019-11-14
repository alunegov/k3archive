package usecase

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	mocks "github.com/alunegov/k3archive/mocks/models"
	"github.com/alunegov/k3archive/models"
)

func TestGetDir(t *testing.T) {
	mapperMock := &mocks.FileDataMapper{}
	mapperMock.On("GetFileInfo", "../../testdata/Rms/2").Return(uint8(2), "c2", nil)
	mapperMock.On("GetFileInfo", "../../testdata/Rms/3").Return(uint8(3), "c3", nil)

	sut := NewArchiveUsecase("../../testdata", "Rms", mapperMock)

	res, err := sut.GetDir()

	if assert.NoError(t, err) {
		ref := []*models.FileInfo{
			{Name: "2", Size: 1264, ModTime: res[0].ModTime, Path: "/Rms/2", Opts: 2, Comment: "c2"},
			{Name: "3", Size: 1264, ModTime: res[1].ModTime, Path: "/Rms/3", Opts: 3, Comment: "c3"},
		}
		assert.Equal(t, ref, res)
	}
}

func TestGetDir_ThenGetFileInfoError(t *testing.T) {
	mapperMock := &mocks.FileDataMapper{}
	mapperMock.On("GetFileInfo", "../../testdata/Rms/2").Return(uint8(0), "", errors.New("expected error"))

	sut := NewArchiveUsecase("../../testdata", "Rms", mapperMock)

	res, err := sut.GetDir()

	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestGetFile(t *testing.T) {
	mapperMock := &mocks.FileDataMapper{}
	mapperMock.On("GetFileInfo", "../../testdata/Rms/2").Return(uint8(2), "c2", nil)

	sut := NewArchiveUsecase("../../testdata", "Rms", mapperMock)

	res, err := sut.GetFile("2")

	if assert.NoError(t, err) {
		ref := &models.FileInfo{Name: "2", Size: 1264, ModTime: res.ModTime, Path: "/Rms/2", Opts: 2, Comment: "c2"}
		assert.Equal(t, ref, res)
	}
}

func TestDeleteFile(t *testing.T) {
	sut := NewArchiveUsecase("../../testdata", "Rms", nil)

	const testFileName = "../../testdata/Rms/99"

	f, err := os.Create(testFileName)
	if !assert.NoError(t, err) {
		return
	}
	_ = f.Close()

	err = sut.DeleteFile("99")
	if !assert.NoError(t, err) {
		return
	}

	_, err = os.Stat(testFileName)
	assert.False(t, os.IsExist(err))
}

func TestDeleteFile_ThenFileNotExists(t *testing.T) {
	sut := NewArchiveUsecase("../../testdata", "Rms", nil)

	err := sut.DeleteFile("0")
	assert.Error(t, err)
}
