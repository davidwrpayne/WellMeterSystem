package service

import (
	"fmt"
	"github.com/davidwrpayne/wellmetersystem/repository"
	"github.com/davidwrpayne/wellmetersystem/schema"
	"github.com/davidwrpayne/wellmetersystem/sensor"
	"github.com/davidwrpayne/wellmetersystem/sor"
	"github.com/google/uuid"
	"time"
)

type WellMeasurement struct {
	store                 repository.Storage
	sensor                sensor.DistanceSensor
	systemOfRecordService sor.Service
}

func NewWellMeasurement(store repository.Storage, sensor sensor.DistanceSensor, recordService sor.Service) *WellMeasurement {
	return &WellMeasurement{
		store:                 store,
		sensor:                sensor,
		systemOfRecordService: recordService,
	}
}

func (s *WellMeasurement) MeasureWell() error {
	distanceCM, err := s.sensor.MeasureCM()
	if err != nil {
		return err
	}
	measuredAt := time.Now()
	measurement := schema.NewMeasurement(uuid.NewString(), distanceCM, measuredAt)
	fmt.Printf("took measurement %v", measurement)
	return s.store.WriteMeasurement(measurement) // process only writes measurement to local storage.
}

func (s *WellMeasurement) ReportAllMeasurements() error {
	unpublished, err := s.store.ReadUnpublished()
	if err != nil {
		return err
	}
	for _, measurement := range unpublished {
		err := s.systemOfRecordService.Publish(measurement)
		if err != nil {
			return err
		} else {
			s.store.MarkPublished(measurement.UUID)
		}
	}
	return nil
}
