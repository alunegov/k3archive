package mocks

import (
	"os"
	"time"

	mock "github.com/stretchr/testify/mock"
)

type FileInfo struct {
	mock.Mock
}

func (it *FileInfo) Name() string {
	args := it.Called()
	return args.String(0)
}

func (it *FileInfo) Mode() os.FileMode {
	args := it.Called()
	return args.Get(0).(os.FileMode)
}

func (it *FileInfo) Size() int64 {
	args := it.Called()
	return args.Get(0).(int64)
}

func (it *FileInfo) ModTime() time.Time {
	args := it.Called()
	return args.Get(0).(time.Time)
}

func (it *FileInfo) IsDir() bool {
	args := it.Called()
	return args.Bool(0)
}

func (it *FileInfo) Sys() interface{} {
	args := it.Called()
	return args.Get(0).(interface{})
}
