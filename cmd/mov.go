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
	RunE: func(cmd *cobra.Command, args []string) error {
		// fmt.Println("mov called", args)
		fmt.Printf("-i: %v, -b: %v, -o: %v, -r: %v\n", flagSearchIterative, flagKeepBoth, flagKeepOld, flagReplaceOld)
		searchMethod := DetermineSearchMethod()
		replaceMethod, err := DetermineFileHandling()
		fmt.Printf("mov called with following flags:\n\tsearchmethod:\t%v\n\treplacemethod:\t%v\n\terror:\t%v\n", searchMethod, replaceMethod, err)
		return err
	},
}

func init() {
	rootCmd.AddCommand(movCmd)
}

func ConvertMovFileToMp4(path string, filename string, originType string, targetType string) bool {
	// convert mov-file to mp4

	return true
}
