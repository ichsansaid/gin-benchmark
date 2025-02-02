package server

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Server started")
	},
}

func Execute() {
	if err := ServerCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
