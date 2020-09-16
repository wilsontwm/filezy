package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var isRecursive bool
var folder string
var prefix string
var suffix string
var regexPattern string
var extension string
var enableLog bool

var RootCmd = &cobra.Command{
	Use:     "filezy",
	Short:   "filezy is a CLI-based file management tool",
	Version: "0.0.1",
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&isRecursive, "recursive", "r", false, "Scan files in sub-directories recursively")
	RootCmd.PersistentFlags().StringVarP(&folder, "folder", "f", "./", "Target folder to be scanned")
	RootCmd.PersistentFlags().StringVarP(&prefix, "prefix", "p", "", "Return files that have the specified prefix in the file name")
	RootCmd.PersistentFlags().StringVarP(&suffix, "suffix", "s", "", "Return files that have the specified suffix in the file name")
	RootCmd.PersistentFlags().StringVarP(&regexPattern, "regex", "x", "", "Return files that match the regex pattern in the file name")
	RootCmd.PersistentFlags().StringVarP(&extension, "ext", "e", "", "Return files that have the specified extension")
	RootCmd.PersistentFlags().BoolVarP(&enableLog, "log", "l", false, "Print logs")
}
func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
