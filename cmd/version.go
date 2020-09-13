package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/wilsontwm/filezy/constant"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of filezy",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(constant.VERSION)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
