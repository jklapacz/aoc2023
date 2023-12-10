package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "aoc",
	Short: "AOC is a CLI for interacting with the Advent of Code.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

func init() {
	switch os.Getenv("LOG_LEVEL") {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	}
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(day01Cmd)
	RootCmd.AddCommand(day02Cmd)
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
