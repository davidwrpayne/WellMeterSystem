package repository

import (
	"errors"
	"fmt"
	"github.com/davidwrpayne/wellmetersystem/schema"
)

type Memory struct {
	data        map[string]*schema.Measurement
	reportedIds []string
}

func NewMemory() *Memory {
	return &Memory{
		data:        map[string]*schema.Measurement{},
		reportedIds: []string{},
	}
}

func (m Memory) FetchUnreported() ([]*schema.Measurement, error) {
	unreported := []*schema.Measurement{}
	for _, id := range m.reportedIds {
		if measurement, ok := m.data[id]; ok {
			unreported = append(unreported, measurement)
		} else {
			return nil, errors.New(fmt.Sprintf("Error: Id %s was not found in data list", id))
		}
	}
	return unreported, nil
}

func (m Memory) Store(measurement *schema.Measurement) error {
	m.data[measurement.UUID] = measurement
	return nil
}

func (m Memory) Load(uuid string) (*schema.Measurement, error) {
	if value, found := m.data[uuid]; found {
		return value, nil
	} else {
		return nil, errors.New(fmt.Sprintf("Load: Could not find measurement for uuid: %s",uuid))
	}
}

func (m Memory) MarkReported(uuid string) error {
	if _, found := m.data[uuid]; found {
		m.reportedIds = append(m.reportedIds, uuid)
		return nil
	} else {
		return errors.New(fmt.Sprintf("MarkReported: Could not find measurement for uuid: %s",uuid))
	}
}

var _ Storage = (*Memory)(nil)
