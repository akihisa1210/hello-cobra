package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	echoCmd.Flags().BoolP("no-linefeed", "n", false, "no line feed")

	rootCmd.AddCommand(echoCmd)
}

var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "Echo arguments",
	Run: func(cmd *cobra.Command, args []string) {
		nolf, _ := cmd.Flags().GetBool("no-linefeed")
		for i, arg := range args {
			fmt.Print(arg)
			if i != len(args)-1 {
				fmt.Print(" ")
			} else {
				if !nolf {
					fmt.Print("\n")
				}
			}
		}
	},
}
