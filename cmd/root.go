package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var definitionFileRootFlag string

var rootCmd = &cobra.Command{
	Use:   "envvars",
	Short: "envvars gives your environment variables the love they deserve",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&definitionFileRootFlag, "file", "f", "envvars.toml", "definition file")
	rootCmd.AddCommand(ensureCmd)
}
