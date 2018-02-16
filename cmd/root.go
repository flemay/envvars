package cmd

import (
	// "fmt"
	"github.com/spf13/cobra"
	"os"
)

var definitionFileRootFlag string
var tagsRootFlag []string

var rootCmd = &cobra.Command{
	Use:   "envvars",
	Short: "envvars gives your environment variables the love they deserve",
	Long: `A way to define environment variables and ensure they comply
Usage examples
  validate the definition file if it contains errors
    $ envvars validate
  ensure the environment variables comply with the definition file
    $ envvars ensure`,
	SilenceUsage: true,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&definitionFileRootFlag, "file", "f", "envvars.toml", "definition file")
	rootCmd.PersistentFlags().StringSliceVarP(&tagsRootFlag, "tags", "t", nil, "execute subcommands against environment variables with the tags (ex: --tags test,build)")
}
