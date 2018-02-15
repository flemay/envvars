package cmd

import (
	"github.com/flemay/envvars/pkg/envvars"
	"github.com/spf13/cobra"
)

var ensureCmd = &cobra.Command{
	Use:   "ensure",
	Short: "Verify that your environment variables comply to their definition",
	RunE: func(cmd *cobra.Command, args []string) error {
		definition, err := envvars.NewDefinition(definitionFileRootFlag)
		if err != nil {
			return err
		}
		return envvars.Ensure(definition, tagsRootFlag...)
	},
}

func init() {
	rootCmd.AddCommand(ensureCmd)
}
