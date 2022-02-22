/*
Copyright © 2022 Brain4Tech <brain4techyt@gmail.com>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
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
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("iterative", "i", false, "Whether the tool should only convert files in current directory instead of converting files recursivly")
	rootCmd.PersistentFlags().BoolP("both", "b", false, `Put both original and converted files in same directory.
Cannot be used together with --k and --r.`)
	rootCmd.PersistentFlags().BoolP("keep", "k", false, `Keep original files and put converted files in subdirectories.
Cannot be used together with --b and --r.`)
	rootCmd.PersistentFlags().BoolP("replace", "r", true, `Replace converted files with original files and backup original files in subdirectories.
Cannot be used together with --k and --b.`)
}
