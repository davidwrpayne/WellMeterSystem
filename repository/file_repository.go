package repository

import (
	"fmt"
	"github.com/davidwrpayne/wellmetersystem/schema"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"os"
)

type FileRepository struct {
	baseDirectory string
	filename      string
}

func (r *FileRepository) ReadUnpublished() (map[string]*schema.Measurement, error) {
	files, err := ioutil.ReadDir(r.unpublishedDirectory())
	if err != nil {
		return nil, err
	}
	result := map[string]*schema.Measurement{}
	for _, f := range files {
		f.Name()
		filePath := fmt.Sprintf("%s/%s", r.unpublishedDirectory(), f.Name())
		data, err := ioutil.ReadFile(filePath)
		if err != nil {
			return nil, err
		}
		newMeasurement := &schema.Measurement{}
		err = yaml.Unmarshal(data, newMeasurement)
		if err != nil {
			return nil, err
		}
		result[newMeasurement.UUID] = newMeasurement
	}
	return result, nil
}

func (r *FileRepository) MarkPublished(uuid string) error {
	oldFilePath := fmt.Sprintf("%s/%s.yaml", r.unpublishedDirectory(), uuid)
	newFilePath := fmt.Sprintf("%s/%s.yaml", r.publishedDirectory(), uuid)
	err := os.Rename(oldFilePath, newFilePath)
	if err != nil {
		return err
	}
	return nil
}

func (r *FileRepository) WriteMeasurement(measurement *schema.Measurement) error {
	fileName := fmt.Sprintf("%s.yaml", measurement.UUID)
	filePath := fmt.Sprintf("%s/%s", r.unpublishedDirectory(), fileName)
	marshaledData, err := yaml.Marshal(measurement)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, marshaledData, 0777)
}

func (r *FileRepository) publishedDirectory() string {
	return fmt.Sprintf("%s/published", r.baseDirectory)
}
func (r *FileRepository) unpublishedDirectory() string {
	return fmt.Sprintf("%s/unpublished", r.baseDirectory)
}

func NewFileRepository(baseDirectory string) *FileRepository {
	return &FileRepository{
		baseDirectory: baseDirectory,
		filename:      "file_repository.yaml",
	}
}

var _ Storage = (*FileRepository)(nil)
