/*
Copyright © 2022 Brain4Tech <brain4techyt@gmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// mp4Cmd represents the mp4 command
var mp4Cmd = &cobra.Command{
	Use:   "mp4",
	Short: "Select all .mp4-files and convert them into .mov",
	Long:  `Select all .mp4-files and convert them into .mov`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mp4 called")
	},
}

func init() {
	rootCmd.AddCommand(mp4Cmd)
}
