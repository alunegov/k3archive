package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	mocks "github.com/alunegov/k3archive/mocks/os"
)

func TestNewFileInfo(t *testing.T) {
	ref := &FileInfo{
		Name:    "name",
		Size:    13,
		ModTime: time.Now(),
		Path:    "test/name",
		Opts:    0,
		Comment: "",
	}

	sysFileInfoMock := &mocks.FileInfo{}
	sysFileInfoMock.On("Name").Return(ref.Name)
	sysFileInfoMock.On("Size").Return(ref.Size)
	sysFileInfoMock.On("ModTime").Return(ref.ModTime)

	res := NewFileInfo(sysFileInfoMock, "test")

	assert.Equal(t, ref, res)
}
