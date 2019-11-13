package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	mocks "github.com/alunegov/k3archive/mocks/os"
)

func TestNewFileInfo(t *testing.T) {
	refTime := time.Now()

	sysFileInfoMock := &mocks.FileInfo{}
	sysFileInfoMock.On("Name").Return("name")
	sysFileInfoMock.On("Size").Return(int64(13))
	sysFileInfoMock.On("ModTime").Return(refTime)

	res := NewFileInfo(sysFileInfoMock, "test")

	assert.Equal(t, "name", res.Name)
	assert.Equal(t, int64(13), res.Size)
	assert.Equal(t, refTime, res.ModTime)
	assert.Equal(t, "test/name", res.Path)
	assert.Equal(t, uint8(0), res.Opts)
	assert.Equal(t, "", res.Comment)
}
