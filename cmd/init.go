package cmd

import (
	"github.com/flemay/envvars/pkg/envvars"
	"github.com/flemay/envvars/pkg/yml"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Creates a new Declaration File",
	Long:  "Creates a new Declaration File to get started",
	RunE: func(cmd *cobra.Command, args []string) error {
		reader := yml.NewDeclarationYML(declarationFileRootFlag)
		return envvars.Init(reader)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
