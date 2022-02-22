/*
Copyright © 2022 Brain4Tech <brain4techyt@gmail.com>

*/
package cmd

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var (
	flagSearchIterative    bool
	flagKeepBoth           bool
	flagKeepOld            bool
	flagReplaceOld         bool
	flagCapitalizeFileType bool
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
Cannot be used together with --o and --r.`)
	rootCmd.PersistentFlags().BoolVarP(&flagKeepOld, "keep-old", "o", false, `Keep original files and put converted files in subdirectories.
Cannot be used together with --b and --r.`)
	rootCmd.PersistentFlags().BoolVarP(&flagReplaceOld, "replace-old", "r", false, `Replace converted files with original files and backup original files in subdirectories (default).
Cannot be used together with --o and --b.`)

	rootCmd.PersistentFlags().BoolVarP(&flagCapitalizeFileType, "uppercase", "u", false, `Whether to search for uppercased file types instead of regular ones (eg. '.MP4' instead of '.mp4'`)

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
	if flagSearchIterative {
		return -1
	}

	return 1
}

func DetermineFileHandling() (int, error) {
	// check flag values to determine how/where to write converted files
	// returns 1 (replace), 2 (keep-both), 3 (keep-old)

	if flagReplaceOld {
		if flagKeepOld {
			if flagKeepBoth {
				return 0, errors.New("too many flags")
			} else {
				return 0, errors.New("--replace-old and --keep-old cannot be combined")
			}
		} else {
			if flagKeepBoth {
				return 0, errors.New("--replace-old and -keep-both cannot be combined")
			} else {
				return 1, nil
			}
		}
	} else {
		if flagKeepOld {
			if flagKeepBoth {
				return 0, errors.New("--keep-old and --keep-both cannot be combined")
			} else {
				return 3, nil
			}
		} else {
			if flagKeepBoth {
				return 2, nil
			} else {
				return 1, nil
			}
		}
	}

}

func DetermineSearchRootPath(args []string) (string, error) {
	if len(args) != 0 {
		// Abs converts any given (absolute) path into a relative one
		return filepath.Abs(args[0])
	}
	return os.Getwd()
}

func SearchFiles(method int, filetype string) ([]File, error) {

	fileArray := make([]File, 0)
	var globMethodString string
	localFileType := filetype
	if flagCapitalizeFileType {
		localFileType = strings.ToUpper(filetype)
	}

	log.Println("Searching for all files ending on ." + localFileType)

	if method < 0 {
		globMethodString = "*." + localFileType

	} else {
		globMethodString = string(os.PathSeparator) + "**" + string(os.PathSeparator) + "*." + localFileType
	}
	foundFiles, err := filepath.Glob(globMethodString)

	if err != nil {
		return make([]File, 0), err
	}

	for _, filepath := range foundFiles {
		file, err := ConvertFilepathToFile(filepath)
		if err == nil {
			fileArray = append(fileArray, file)
		}
	}

	return fileArray, nil
}

func ConvertFilepathToFile(path string) (File, error) {
	// searching files retrieves a path
	// exclude data from this path and store them in a struct for easier file conversion

	/*
		1. split path at '/'
		2. take last element (filename) and split at '.'
		3. recombine splittedpath from root to parent directory of file
		4. store everything in new File struct

		* look into filepath documentation for some cool functions to use for above structure
	*/

	return File{}, nil
}
