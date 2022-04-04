package repository

import "github.com/davidwrpayne/wellmetersystem/schema"



type Storage interface {
	WriteMeasurement(measurement *schema.Measurement) error
	ReadUnpublished() (map[string]*schema.Measurement, error)
	MarkPublished(string) error
}


