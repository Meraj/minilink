package main

import (
	"fmt"
	"os"

	"minilink/cmd"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "minilink",
	Short: "A minimal URL shortener",
	Long:  `MiniLink is a simple URL shortener that uses YAML files as a database.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(cmd.GetServeCmd())
}

func main() {
	Execute()
}