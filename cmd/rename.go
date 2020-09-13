package cmd

import (
	"fmt"
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

var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename files in batch",
	Run: func(cmd *cobra.Command, args []string) {
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

		for i, file := range files {
			fmt.Printf("%v: %+v\n", helper.ToString(i+1, helper.NumberOfDigits(total)), file)
		}

		fmt.Println("Total files:", total, "Number of digit:", helper.NumberOfDigits(total))

	},
}

func init() {
	renameCmd.Flags().BoolVarP(&isRecursive, "recursive", "r", false, "Scan files in sub-directories recursively")
	renameCmd.Flags().StringVarP(&folder, "folder", "f", "./", "Target folder to be scanned")
	renameCmd.Flags().StringVarP(&prefix, "prefix", "p", "", "Return files that have the specified prefix in the file name")
	renameCmd.Flags().StringVarP(&suffix, "suffix", "s", "", "Return files that have the specified suffix in the file name")
	renameCmd.Flags().StringVarP(&regexPattern, "regex", "x", "", "Return files that match the regex pattern in the file name")
	renameCmd.Flags().StringVarP(&extension, "ext", "e", "", "Return files that have the specified extension")

	RootCmd.AddCommand(renameCmd)
}
