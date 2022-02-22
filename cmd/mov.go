/*
Copyright © 2022 Brain4Tech <brain4techyt@gmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// movCmd represents the mov command
var movCmd = &cobra.Command{
	Use:   "mov",
	Short: "Select all .mov-files and convert them into .mp4",
	Long:  `Select all .mov-files and convert them into .mp4`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mov called")
	},
}

func init() {
	rootCmd.AddCommand(movCmd)
}
