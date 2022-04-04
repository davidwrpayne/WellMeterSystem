package cli

import (
	"fmt"
	"github.com/davidwrpayne/wellmetersystem/repository"
	"github.com/davidwrpayne/wellmetersystem/service"
	"github.com/spf13/cobra"
)

var reportCmd = &cobra.Command{
	Use: "report",
	Short: "report measurements from unpublished directory.",
	Long: "reports measurements found in the unpublished directory to the System of Record and moves them to the published directory.",
	Run: func(cmd *cobra.Command, args []string) {
		reportMeasurements()
	},
}

func reportMeasurements() {
	service := service.NewWellMeasurement(repository.NewFileRepository("./data"), nil)
	err := service.ReportAllMeasurements()
	if err != nil {
		fmt.Errorf("Error reporting measurements: %s", err)
		return
	}
}



func init() {
	rootCmd.AddCommand(reportCmd)
}