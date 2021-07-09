package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(echoCmd)
}

var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "Echo arguments",
	Run: func(cmd *cobra.Command, args []string) {
		for i, arg := range args {
			fmt.Print(arg)
			if i != len(args)-1 {
				fmt.Print(" ")
			} else {
				fmt.Print("\n")
			}
		}
	},
}
