package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "controller",
	SilenceErrors: true,
	SilenceUsage:  true,
}

func Execute() error {
	// Cancel context with CTRL^C
	ctx, close := context.WithCancel(context.Background())
	sigChan := HandleInterrupt()
	go func() {
		<-sigChan
		fmt.Println("Received interrupt ")
		close()
	}()

	return rootCmd.ExecuteContext(ctx)
}

func HandleInterrupt() chan os.Signal {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	return sigChan
}
