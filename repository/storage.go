package repository

import "github.com/davidwrpayne/wellmetersystem/schema"

type Storage interface {
	FetchUnreported() ([]*schema.Measurement, error)
	Store(measurement *schema.Measurement) error
	Load(uuid string) (*schema.Measurement, error)
	MarkReported(id string) error
}