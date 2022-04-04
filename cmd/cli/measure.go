package cli

import (
	"fmt"
	"github.com/davidwrpayne/wellmetersystem/repository"
	"github.com/davidwrpayne/wellmetersystem/service"
	"github.com/spf13/cobra"
)

var measureCmd = &cobra.Command{
	Use: "measure",
	Short: "measure using configured distance sensor",
	Long: "measure using configured distance sensor.",
	Run: func(cmd *cobra.Command, args []string) {
		measure_distance()
	},
}



func measure_distance() {
	service := service.NewWellMeasurement(repository.NewFileRepository("./data"), nil)
	err := service.MeasureWell()
	if err != nil {
		fmt.Errorf("Error measureing well %s", err)
		return
	}
}


func init() {
	rootCmd.AddCommand(measureCmd)
}