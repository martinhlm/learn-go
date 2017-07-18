package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var CfgFile string

var RootCmd = &cobra.Command{
	Use:   "dagobah",
	Short: `Dagobah is an awesome planet`,
	Long: `Dagobah provides planet style RSS aggregation. It is inspired by
	python planet. Simple yaml config and it's own server`,
	Run: rootRun,
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&CfgFile, "config", "", "config file")
}

func initConfig() {
	if CfgFile != "" {
		viper.SetConfigFile(CfgFile)
	}
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/dagobah/")
	viper.AddConfigPath("$GOPATH/src/testing/first-go-app/.dagobah/")
	viper.ReadInConfig()
}

func Execute() {
	addCommands()
	err := RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func rootRun(cmd *cobra.Command, args []string) {
	fmt.Println(viper.Get("feeds"))
	fmt.Println(viper.GetString("appname"))
}

func addCommands() {
	RootCmd.AddCommand(fetchCmd)
}
