# mp4-to-mov
*mp4-to-mov* is a small handy CLI to (recursively) convert .mp4-files in a given directory into a .mov-file, and vice-versa.

## Overview
mp4-to-mov is a CLI specifically designed for Davinci Resolve users on Linux, as it is not possible to use mp4-encoded files with an `acc`-audio codec.
This tool is meant to quickly convert .mp4-files to a usable .mov-fileformat with `pcm_s16le` audio-encoding.

The same process can be done in reverse, so rendered projects can be shared in an .mp4-format. Converting files can be done recursively in subdirectories or just the given directory. You can choose whether to replace old files with freshly converted ones, or to move old/new files into subdirectories.

## Installation

**1 Check dependencies**
* [Go](https://go.dev/) >= 1.17.6
* [ffmpeg](https://www.ffmpeg.org/)

**2a) Local Installation**

1. Download/Clone this repository
2. `cd` into the repository and type `go install github.com/Brain4Tech/mp4-to-mov`
3. Use it from everywhere with `mp4-to-mov`. `$GOPATH/bin` must be in your PATH in order for it to work

*OR*

**2b) Use pre-compiled binaries (You don't need Go as a dependency)**

* // TODO

## Usage
`mp4-to-mov [command] [--flags]`
   
**Available Commands:**

 - `help`: Help about any command
 - `mov`: Select all .mov-files and convert them into .mp4
 - `mp4`: Select all  mp4-files and convert them into .mov

**Flags:**
* `-h`, `--help`: help for mp4-to-mov
* `-i`, `--iterative`: Whether the tool should only convert files in current directory instead of converting files recursively
* `-b`, `--keep-both`: Put both original and converted files in same directory. Cannot be used together with --o and --r.
* `-o`, `--keep-old`: Keep original files and put converted files in subdirectories. Cannot be used together with --b and --r.
* `-r`, `--replace-old`: Replace converted files with original files and backup original files in subdirectories (default). Cannot be used together with --o and --b.
* `-u`, `--uppercase`: Whether to search for uppercased file types instead of regular ones (eg. '.MP4' instead of '.mp4')

Use `mp4-to-mov [command] --help` for more information about a command.

## About and about me
* This is my first project written in Go. I'm sure some stuff can be optimized and improved.
* [Cobra](https://github.com/spf13/cobra)-library is used to create CLI biolerplate code
* English is not my first language, so expect some typos and grammar mistakes :D
* Developed and compiled on *Ubuntu 20.04.3 LTS* and *go1.17.6*; kernel version *5.13.0-30-generic*

## Contributing
Found a typo? Found a bug or optimization issue?
If you want to contribute, feel free!


(c) 2022 Brain4Tech