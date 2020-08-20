package cmd

import (
	"fmt"

	"C"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var (
	VERSION     string
	BUILD_TIME  string
	GO_VERSION  string
	GIT_BRANCH  string
	COMMIT_SHA1 string
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Bridge",
	Long:  `This is Bridge's version`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("version:\t%s \nbuild time:\t%s\ngit branch:\t%s\ngit commit:\t%s\ngo version:\t%s\n", VERSION, BUILD_TIME, GIT_BRANCH, COMMIT_SHA1, GO_VERSION)
	},
}
