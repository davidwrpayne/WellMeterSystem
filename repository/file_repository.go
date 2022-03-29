package repository

import (
	"github.com/davidwrpayne/wellmetersystem/schema"
)

type FileRepository struct {
	Path string
}

func NewFileRepository(path string) *FileRepository {
	return &FileRepository{
		Path: path,
	}
}

var _ Storage = (*FileRepository)(nil)

func (f FileRepository) Load(uuid string) (*schema.Measurement, error) {
	panic("implement me")
}

func (f FileRepository) FetchUnreported() ([]*schema.Measurement, error) {
	panic("implement me")
}

func (f FileRepository) Store(measurement *schema.Measurement) error {
	panic("implement me")
}

func (f FileRepository) MarkReported(id string) error {
	panic("implement me")
}






