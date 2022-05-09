package cli

import (
	"fmt"
	"github.com/davidwrpayne/wellmetersystem/config"
	"github.com/davidwrpayne/wellmetersystem/service"
	"github.com/spf13/cobra"
)

var measureCmd = &cobra.Command{
	Use:   "measure",
	Short: "measure using configured distance sensor",
	Long:  "measure using configured distance sensor.",
	Run: func(cmd *cobra.Command, args []string) {
		measure_distance()
	},
}

func measure_distance() {

	sensor, err := configureDistanceSensor(config.SystemConfig)
	if err != nil {
		fmt.Errorf("Failed to get a distance sensor because: %s", err)
		return
	}

	systemOfRecord, err := configureSOR(config.SystemConfig)
	if err != nil {
		fmt.Errorf("Failed to configure sor : %s", err)
		return
	}

	storage := configureFileRepository()
	service := service.NewWellMeasurement(storage, sensor, systemOfRecord)
	err = service.MeasureWell()
	if err != nil {
		fmt.Errorf("Error measuring well %s", err)
		return
	}
}

func init() {
	rootCmd.AddCommand(measureCmd)
}
