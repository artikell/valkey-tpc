/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/artikell/valkey-tpc/workload"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// prepareCmd represents the prepare command
var prepareCmd = &cobra.Command{
	Use:   "prepare",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		w := viper.GetString("workload")
		workload.RunWorkLoad(w, workload.ModePrepare)
	},
}

func init() {
	rootCmd.AddCommand(prepareCmd)
}
