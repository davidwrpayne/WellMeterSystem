package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use: "hello",
	Short: "test",
	Long: "test",
	Run: func(cmd *cobra.Command, args []string) {

		helloWorld()
	},
}


func init() {
	rootCmd.AddCommand(helloCmd)
}

func helloWorld() {
	fmt.Printf("Hello World!\n")
}
