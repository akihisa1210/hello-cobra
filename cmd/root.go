package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "hello",
		Short: "A short description",
		Long:  "A long description",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("hello!")
		},
	}
)

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
