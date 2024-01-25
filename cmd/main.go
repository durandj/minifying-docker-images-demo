package main

import (
	"fmt"
	"os"

	"github.com/durandj/minifying-docker-images-demo/internal"
	"github.com/spf13/cobra"
)

func newRootCommand() *cobra.Command {
	return &cobra.Command{
		Use: "demo",
		RunE: func(cmd *cobra.Command, args []string) error {
			server := internal.New()

			if err := server.Start(); err != nil {
				return err
			}

			return nil
		},
	}
}

func main() {
	rootCmd := newRootCommand()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
