/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/viper"
	"golang.org/x/exp/slog"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "valkey-tpc",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		slog.Info(cfgFile)

		viper.SetConfigFile(cfgFile)
		viper.SetConfigType("yaml")

		// Read config file
		if err := viper.ReadInConfig(); err != nil {
			slog.Error("config file not found or another error was produced", err)
			return
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	slog.Info("Start valkey-tpc...")
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	dir, err := os.Getwd()
	if err != nil {
		slog.Error("", err)
	}

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", dir+"/.valkey-tpc.yaml", "config file (default is $PWD/.valkey-tpc.yaml)")
}
