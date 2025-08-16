package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "go-rest-sample",
	Short: "Backend Go rest sample",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	_ = viper.BindEnv("App.Host", "APP_HOST")
	_ = viper.BindEnv("App.Port", "APP_PORT")
}
