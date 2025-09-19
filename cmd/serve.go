package cmd

import (
	"log"

	"minilink/internal/config"
	"minilink/internal/server"

	"github.com/spf13/cobra"
)

var (
	configFile string
	port       string
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the URL shortener server",
	Long:  `Start the HTTP server to handle URL redirects based on the YAML configuration.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig(configFile)
		if err != nil {
			log.Fatalf("Failed to load config: %v", err)
		}

		srv := server.New(cfg)
		if err := srv.Start(port); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	},
}

func init() {
	serveCmd.Flags().StringVarP(&configFile, "config", "c", "links.yaml", "Path to the YAML config file")
	serveCmd.Flags().StringVarP(&port, "port", "p", "8080", "Port to run the server on")
}

func GetServeCmd() *cobra.Command {
	return serveCmd
}