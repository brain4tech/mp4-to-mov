/*
Copyright © 2022 Brain4Tech <brain4techyt@gmail.com>

*/
package cmd

import (
	"log"
	"os/exec"

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
		log.Println("Found", len(fileArray), "files")

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
	conversion, err := exec.Command("ffmpeg", "-i", file.parentPath+"/"+file.Name+file.Type, "-acodec", "pcm_s16le", "-vcodec", "copy", file.parentPath+"/"+file.Name+".mov").Output()
	log.Println("conversion", conversion, err)

	if replaceMethod == 1 {
		newPath := "mkdir " + file.parentPath + "/old"
		_, errPathCreation := exec.Command("mkdir", file.parentPath+"/old").Output()
		log.Println("creating path:", newPath, errPathCreation)

		fileMove, errFileMove := exec.Command("mv", file.parentPath+"/"+file.Name+file.Type, file.parentPath+"/old/"+file.Name+file.Type).Output()
		log.Println("moving old file:", fileMove, errFileMove)

	} else if replaceMethod == 3 {
		_, errNewPath := exec.Command("mkdir", file.parentPath+"/converted").Output()
		log.Println("creating path:", errNewPath)

		_, errFileMove := exec.Command("mv", file.parentPath+"/"+file.Name+".mov", file.parentPath+"/converted/"+file.Name+".mov").Output()
		log.Println("moving converted file:", errFileMove)
	}

	return true
}
