package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/LoftLabs-Experiments/test-release-oss/pkg/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the test application server",
	Long:  `Start the test application server that simulates basic functionality.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting test application server...")

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// Start dummy server
		srv := server.NewTestServer()
		go func() {
			if err := srv.Start(ctx); err != nil {
				logrus.WithError(err).Error("Failed to start server")
				os.Exit(1)
			}
		}()

		// Wait for interrupt signal
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c

		logrus.Info("Shutting down gracefully...")
		cancel()
		time.Sleep(2 * time.Second)
	},
}
