package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"os"
)

var declarationFileRootFlag string
var tagsRootFlag []string

var rootCmd = &cobra.Command{
	Use:   "envvars",
	Short: "Envvars gives your environment variables the love they deserve",
	Long: `A way to declare environment variables and ensure they comply
Usage examples
  validate the declaration file if it contains errors
    $ envvars validate
  ensure the environment variables comply with the declaration file
    $ envvars ensure`,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Help()
		return errors.New("A command needs to be provided to envvars")
	},
}

func Execute(version, commitHash, buildDate string) {
	envvarsVersion = version
	envvarsCommitHash = commitHash
	envvarsBuildDate = buildDate
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&declarationFileRootFlag, "file", "f", "envvars.yml", "declaration file")
	rootCmd.PersistentFlags().StringSliceVarP(&tagsRootFlag, "tags", "t", nil, "list of tags targetting a subset of environment variables (ex: --tags test,build)")
}
