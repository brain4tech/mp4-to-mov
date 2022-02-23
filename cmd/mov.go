/*
Copyright © 2022 Brain4Tech <brain4techyt@gmail.com>

*/
package cmd

import (
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

// movCmd represents the mov command
var movCmd = &cobra.Command{
	Use:   "mov",
	Short: "Select all .mov-files and convert them into .mp4",
	Long:  `Select all .mov-files and convert them into .mp4`,
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
		log.Println("Searching root path:", searchRootPath)

		fileArray := SearchFiles(searchMethod, "mov")
		log.Println(len(fileArray), "files found")

		for _, file := range fileArray {
			success := ConvertMovFileToMp4(file, replaceMethod)
			if !success {
				message := "\tError while trying to convert " + file.parentPath + "/" + file.Name + file.Type
				log.Println(message)
			}
		}
		return err
	},
}

func init() {
	rootCmd.AddCommand(movCmd)
}

func ConvertMovFileToMp4(file File, replaceMethod int) bool {
	// convert mp4-file to mov
	log.Println("Converting", file.parentPath+"/"+file.Name+file.Type)
	_, err := exec.Command("ffmpeg", "-i", file.parentPath+"/"+file.Name+file.Type, "-acodec", "mp3", "-vcodec", "copy", file.parentPath+"/"+file.Name+".mp4").Output()
	if err != nil {
		return false
	}

	if replaceMethod == 1 {
		_, errPathCreation := exec.Command("mkdir", file.parentPath+"/old").Output()
		if errPathCreation != nil {
			log.Println("\tError while creating path", file.parentPath+"/old")
		} else {
			_, errFileMove := exec.Command("mv", file.parentPath+"/"+file.Name+file.Type, file.parentPath+"/old/"+file.Name+file.Type).Output()
			if errFileMove != nil {
				log.Println("\tError while moving", file.parentPath+file.Name+file.Type, "to", file.parentPath+"/old/"+file.Name+file.Type)
			}

		}

	} else if replaceMethod == 3 {
		_, errNewPath := exec.Command("mkdir", file.parentPath+"/converted").Output()
		if errNewPath != nil {
			log.Println("\tError while creating path", file.parentPath+"/converted")
		} else {
			_, errFileMove := exec.Command("mv", file.parentPath+"/"+file.Name+".mp4", file.parentPath+"/converted/"+file.Name+".mp4").Output()
			if errFileMove != nil {
				log.Println("\tError while moving", file.parentPath+file.Name+".mp4", "to", file.parentPath+"/converted/"+file.Name+".mp4")
			}
		}

	}

	return true
}
