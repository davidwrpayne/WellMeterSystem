package cli

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

// rootCommand represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "well-meter-system",
	Short: "Well Meter System",
	Long: `Well Meter System provides a cli to record and report sensor measurements`,
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	//home, err := homedir.Dir()
	//cobra.CheckErr(err)
	//viper.AddConfigPath(home)
	//viper.SetConfigName(".cli")

	viper.AutomaticEnv()
	//if err := viper.ReadInConfig(); err == nil {
	//	fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	//}
}