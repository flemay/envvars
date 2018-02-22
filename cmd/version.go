package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
)

var (
	envvarsVersion    string
	envvarsBuildDate  string
	envvarsCommitHash string
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the envvars version information",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf(`envvars:
version     : %s
build date  : %s
git hash    : %s
go version  : %s
go compiler : %s
platform    : %s/%s
`, envvarsVersion, envvarsBuildDate, envvarsCommitHash, runtime.Version(), runtime.Compiler, runtime.GOOS, runtime.GOARCH)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
