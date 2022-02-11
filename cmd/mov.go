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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// movCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// movCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
