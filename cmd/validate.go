package cmd

import (
	"github.com/flemay/envvars/pkg/envvars"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Check if the definition file contains any error",
	Long:  "The flat tags has no effect with this command",
	RunE: func(cmd *cobra.Command, args []string) error {
		definition, err := envvars.NewDefinition(definitionFileRootFlag)
		if err != nil {
			return err
		}
		return envvars.Validate(definition)
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
