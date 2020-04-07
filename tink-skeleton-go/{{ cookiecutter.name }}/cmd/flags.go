package cmd

import (
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func registerStringFlag(cmd *cobra.Command, name, defaultValue, desc string) {
	var _tmpVar string
	// TODO: Could we use String instead of StringVar here to avoid using _tmpVar?
	cmd.PersistentFlags().StringVar(&_tmpVar, name, defaultValue, desc)
	if err := viper.BindPFlag(name, cmd.PersistentFlags().Lookup(name)); err != nil {
		log.Fatalln(err)
	}
}

func registerDurationFlag(cmd *cobra.Command, name string, defaultValue time.Duration, desc string) {
	var _tmpVar time.Duration
	cmd.PersistentFlags().DurationVar(&_tmpVar, name, defaultValue, desc)
	if err := viper.BindPFlag(name, cmd.PersistentFlags().Lookup(name)); err != nil {
		log.Fatalln(err)
	}
}

func getString(key string, required bool) string {
	if required && viper.GetString(key) == "" {
		log.Fatalf("'%s' configuration parameter not set.", key)
	}
	return viper.GetString(key)
}

func getDuration(key string, required bool) time.Duration {
	if required && viper.GetDuration(key) == 0 {
		log.Fatalf("'%s' configuration parameter not set.", key)
	}
	return viper.GetDuration(key)
}
