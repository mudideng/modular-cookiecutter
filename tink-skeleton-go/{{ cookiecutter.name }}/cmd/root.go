// Package cmd contains the CLI commands for this application.
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "{{ cookiecutter.name }}",
	Short: "{{ cookiecutter.name }} is an example skeleton Go application.",
	Long:  "{{ cookiecutter.name }} is an example microservice written in Go, intended to contain best practices.",
}

var cfgFile string

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config-file", "", "YAML (*.yaml) config file.")
}

func initConfig() {
	if cfgFile == "" {
		return
	}

	// Use config file from the flag.
	viper.SetConfigFile(cfgFile)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}

// Execute handles all CLI flag parsing logic and delegates to relevant
// registered subcommand for actual execution.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
