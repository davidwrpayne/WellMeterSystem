package cli

import (
	"fmt"
	"github.com/davidwrpayne/wellmetersystem/config"
	"github.com/davidwrpayne/wellmetersystem/service"
	"github.com/spf13/cobra"
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "report measurements from unpublished directory.",
	Long:  "reports measurements found in the unpublished directory to the System of Record and moves them to the published directory.",
	Run: func(cmd *cobra.Command, args []string) {
		reportMeasurements()
	},
}

func reportMeasurements() {

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
	err = service.ReportAllMeasurements()
	if err != nil {
		fmt.Errorf("Error reporting measurements: %s", err)
		return
	}
}

func init() {
	rootCmd.AddCommand(reportCmd)
}
