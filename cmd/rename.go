package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

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

var renameCmd = &cobra.Command{
	Use:   "rename [filename]",
	Short: "Rename files in batch, auto-increment number will be added as suffix",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
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
			must(err)
		}

		total := len(files)
		noOfDigits := helper.NumberOfDigits(total)
		for i, file := range files {
			newFileName := file.Folder + filename + file.Ext
			numberStr := helper.ToString(i+1, noOfDigits)
			if total > 1 {
				newFileName = file.Folder + filename + "-" + numberStr + file.Ext
			}

			err := os.Rename(file.FullPath, newFileName)
			must(err)

			if enableLog {
				fmt.Printf("%v: Rename %v --> %v\n", time.Now().Format(time.RFC3339), file.FullPath, newFileName)
			}
		}
	},
}

func init() {
	renameCmd.Flags().BoolVarP(&isRecursive, "recursive", "r", false, "Scan files in sub-directories recursively")
	renameCmd.Flags().StringVarP(&folder, "folder", "f", "./", "Target folder to be scanned")
	renameCmd.Flags().StringVarP(&prefix, "prefix", "p", "", "Return files that have the specified prefix in the file name")
	renameCmd.Flags().StringVarP(&suffix, "suffix", "s", "", "Return files that have the specified suffix in the file name")
	renameCmd.Flags().StringVarP(&regexPattern, "regex", "x", "", "Return files that match the regex pattern in the file name")
	renameCmd.Flags().StringVarP(&extension, "ext", "e", "", "Return files that have the specified extension")
	renameCmd.Flags().BoolVarP(&enableLog, "log", "l", false, "Print logs")

	RootCmd.AddCommand(renameCmd)
}
