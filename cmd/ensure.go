package cmd

import (
	"github.com/flemay/envvars/pkg/envvars"
	"github.com/flemay/envvars/pkg/yml"
	"github.com/spf13/cobra"
)

var ensureCmd = &cobra.Command{
	Use:   "ensure",
	Short: "Verify that the environment variables comply to their declaration",
	RunE: func(cmd *cobra.Command, args []string) error {
		reader := yml.NewDeclarationYML(declarationFileRootFlag)
		return envvars.Ensure(reader, tagsRootFlag...)
	},
}

func init() {
	rootCmd.AddCommand(ensureCmd)
}
