package cli

import (
	"errors"
	"github.com/davidwrpayne/wellmetersystem/config"
	"github.com/davidwrpayne/wellmetersystem/repository"
	"github.com/davidwrpayne/wellmetersystem/sensor"
	"github.com/davidwrpayne/wellmetersystem/sor"
)

func configureDistanceSensor(cfg *config.SystemConfiguration) (sensor.DistanceSensor, error) {
	switch cfg.Sensor {
	case config.FakeSensor:
		return sensor.NewFakeSensor([]float64{1.0, 2.0, 3.0}), nil
	case config.RaspberryPi:
		bcmMode := 1
		return sensor.NewRPIDistanceSensor(cfg.RpiSensorConfig.GpioTriggerPin, cfg.RpiSensorConfig.GpioEchoPin, bcmMode, false), nil
	default:
		return nil, errors.New("Invalid sensor config value")
	}
}

func configureSOR(systemConfig *config.SystemConfiguration) (sor.Service, error) {
	var systemOfRecord sor.Service
	if config.SystemConfig.DisableReporting {
		systemOfRecord = sor.NewLoggingSystemOfRecord(systemConfig.RepositoryHost, systemConfig.RepositoryToken)
	} else {
		systemOfRecord = sor.NewHttpSystemOfRecord(systemConfig.RepositoryHost, systemConfig.RepositoryToken)
	}

	return systemOfRecord, nil
}

func configureFileRepository() repository.Storage {
	return repository.NewFileRepository(config.SystemConfig.RepositoryFolderPath)
}
