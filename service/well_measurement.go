package service

import (
	"github.com/davidwrpayne/wellmetersystem/client"
	"github.com/davidwrpayne/wellmetersystem/repository"
	"github.com/davidwrpayne/wellmetersystem/schema"
	"github.com/davidwrpayne/wellmetersystem/sensor"
	"github.com/google/uuid"
	"time"
)

type WellMeasurement struct {
	store repository.Storage
	sensor    sensor.DistanceSensor
	sorClient client.SystemOfRecord
}


func NewWellMeasurement(store repository.Storage, sensor sensor.DistanceSensor) *WellMeasurement {
	return &WellMeasurement{
		store: store,
		sensor: sensor,
	}
}


func (s *WellMeasurement) MeasureWell() error {
	distanceCM, err := s.sensor.MeasureCM()
	if err != nil {
		return err
	}
	measuredAt := time.Now()
	measurement := schema.NewMeasurement(uuid.NewString(), distanceCM, measuredAt)
	return s.store.WriteMeasurement(measurement) // process only writes measurement to local storage.
}

func (s *WellMeasurement) ReportAllMeasurements() error {
	unpublished, err := s.store.ReadUnpublished()
	if err != nil {
		return err
	}
	for _, measurement := range unpublished {
		err := s.sorClient.Publish(measurement)
		if err != nil {
			return err
		} else {
			s.store.MarkPublished(measurement.UUID)
		}
	}
	return nil
}