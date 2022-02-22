/*
Copyright © 2022 Brain4Tech <brain4techyt@gmail.com>

*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	flagSearchIterative bool
	flagKeepBoth        bool
	flagKeepOld         bool
	flagReplaceOld      bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mp4-to-mov",
	Short: "A CLI to quickly convert .mp4 to .mov-files and vice-versa",
	Long: `mp4-to-mov is a CLI specifically designed for Davinci Resolve users on Linux,
as it is not possible to use mp4-encoded files with an acc-audio codec.
This tool is meant to quickly convert .mp4-files to a usable .mov-fileformat
with pcm_s16le audio-encoding.

The same process can be made in reverse, so rendered projects can be shared
in an .mp4-format. Converting files can be done recursivly in
subdirectories or just the given directory.
Converted files can be replaced or stored in a new subdirectory.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&flagSearchIterative, "iterative", "i", false, "Whether the tool should only convert files in current directory instead of converting files recursivly")
	rootCmd.PersistentFlags().BoolVarP(&flagKeepBoth, "keep-both", "b", false, `Put both original and converted files in same directory.
Cannot be used together with --k and --r.`)
	rootCmd.PersistentFlags().BoolVarP(&flagKeepOld, "keep-old", "k", false, `Keep original files and put converted files in subdirectories.
Cannot be used together with --b and --r.`)
	rootCmd.PersistentFlags().BoolVarP(&flagReplaceOld, "replace-old", "r", true, `Replace converted files with original files and backup original files in subdirectories.
Cannot be used together with --k and --b.`)
}

type File struct {
	parentPath string
	Name       string
	Type       string
	ContraType string
}

func DetermineSearchMethod() int {
	// check flag variables to determine whether to search files recursivly or just the directory
	// returns 1 (recursive, as it is default) or -1 (given directory)

	return 0
}

func DetermineFileHandling() int {
	// check flag values to determine how/where to write converted files
	// returns 1 (replace), 2 (keep-both), 3 (keep-old)

	return 0
}

func SearchFiles(method int) []File {
	// search all files according to the method
	// store them in an array of strings

	return make([]File, 0)
}

func ConvertFilepathToFile(path string) File {
	// searching files retrieves a path
	// exclude data from this path and store them in a struct for easier file conversion

	return File{}
}
