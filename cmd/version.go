package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	VERSION string
	COMMIT  string
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Hugo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
		fmt.Printf("Version: %s, Commit: %s \n", VERSION, COMMIT)
	},
}
