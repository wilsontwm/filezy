package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/wilsontwm/filezy/helper"
)

var renameCmd = &cobra.Command{
	Use:   "rename [filename]",
	Short: "Rename files in batch, auto-increment number will be added as suffix",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]

		files, err := getFilteredFiles(folder)
		must(err)

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
	renameCmd.Flags().StringVarP(&folder, "folder", "f", "./", "Target folder to be scanned")

	RootCmd.AddCommand(renameCmd)
}
