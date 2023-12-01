package cmd

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of AOC",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("AOC 2023")
	},
}
