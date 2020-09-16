package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"

	"github.com/wilsontwm/filezy/helper"
	"github.com/wilsontwm/filezy/model"
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

// Get the filtered files
func getFilteredFiles(folder string) ([]model.File, error) {
	files := make([]model.File, 0)
	if filesInFolder, err := helper.GetFiles(folder, isRecursive); err == nil {
		for _, file := range filesInFolder {
			if prefix != "" && !strings.HasPrefix(file.FileName, prefix) {
				continue
			} else if suffix != "" && !strings.HasSuffix(file.FileName, suffix) {
				continue
			} else if regexPattern != "" && !regexp.MustCompile(regexPattern).MatchString(file.FileName) {
				continue
			} else if extension != "" && strings.TrimSuffix(file.Ext, extension) != "." {
				continue
			}

			files = append(files, file)
		}
	} else {
		return files, err
	}

	return files, nil
}
