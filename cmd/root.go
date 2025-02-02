package cmd

import (
	"fmt"
	"os"

	"github.com/ichsansaid/golang-gin-benchmark/cmd/server"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "benchmark",
	Short: "A benchmark using GoRoutines",
}

func Execute() {
	rootCmd.AddCommand(server.ServerCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
