package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
)

// with values that work in a local development environment.
var SystemConfig *SystemConfiguration

func init() {
	SystemConfig = &SystemConfiguration{}
	err := LoadConfigFromEnv(SystemConfig)
	if err != nil {
		panic(fmt.Sprintf("loading config file: %s", err))
	}
}

type SensorType string

const (
	FakeSensor    SensorType = "Fake"
	RaspberryPi   SensorType = "RaspberryPiSensor"
	InvalidSensor SensorType = "Invalid"
)

type RaspberryPiSensorConfig struct {
	GpioTriggerPin int `default:"17" split_words:"true"`
	GpioEchoPin    int `default:"22" split_words:"true"`
}

type SystemConfiguration struct {
	Sensor               SensorType              `default:"Fake"`
	RpiSensorConfig      RaspberryPiSensorConfig `split_words:"true"`
	ReportMeasurement    bool                    `default:"false"`
	RepositoryHost       string                  `envconfig:"WATER_TANK_MONITOR_SERVICE_URL"`
	RepositoryToken      string                  `envconfig:"MEASUREMENT_UPLOAD_API_BEARER_TOKEN"`
	RepositoryFolderPath string                  `envconfig:"MEASUREMENT_STORAGE_PATH"`
}

func LoadConfigFromEnv(cfg *SystemConfiguration) error {

	err := envconfig.Process("wms", cfg)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}
