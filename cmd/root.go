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
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mp4-2-mov.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
