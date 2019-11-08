package archive

import (
	"github.com/alunegov/k3archive/models"
)

type Usecase interface {
	GetDir() ([]*models.FileInfo, error)
	GetFile(id string) (*models.FileInfo, error)
	DeleteFile(id string) error
}
