/*
Copyright © 2022 Brain4Tech <brain4techyt@gmail.com>

*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// mp4Cmd represents the mp4 command
var mp4Cmd = &cobra.Command{
	Use:   "mp4",
	Short: "Select all .mp4-files and convert them into .mov",
	Long:  `Select all .mp4-files and convert them into .mov`,
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// fmt.Println("mp4 called")
		searchMethod := DetermineSearchMethod()
		replaceMethod, err := DetermineFileHandling()
		if err != nil {
			return err
		}
		log.Println("searchMethod:", searchMethod, "|", "replaceMethod:", replaceMethod)

		searchRootPath, err := DetermineSearchRootPath(args)
		if err != nil {
			return err
		}
		log.Println("Using", searchRootPath, "as root path for search")

		fileArray, err := SearchFiles(searchMethod, "mp4")
		if err != nil {
			return err
		}

		for _, file := range fileArray {
			success := ConvertMp4FileToMov(file, replaceMethod)
			if !success {
				message := "Error while trying to convert " + file.parentPath + file.Name + "." + file.Type
				log.Fatal(message)
			}
		}
		return err
	},
}

func init() {
	rootCmd.AddCommand(mp4Cmd)
}

func ConvertMp4FileToMov(file File, replaceMethod int) bool {
	// convert mp4-file to mov

	if replaceMethod == 1 {
		// create subdirectory 'old'
		// convert file
		// move old file to 'old'
	} else if replaceMethod == 2 {
		// convert files
	} else {
		// create subdirectory 'converted'
		// directly convert file into 'converted'
	}

	return true
}
