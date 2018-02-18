package cmd

import (
	"github.com/flemay/envvars/pkg/envvars"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Check if the declaration file contains any error",
	Long:  "The flag tags has no effect with this command",
	RunE: func(cmd *cobra.Command, args []string) error {
		d, err := envvars.NewDeclaration(declarationFileRootFlag)
		if err != nil {
			return err
		}
		return envvars.Validate(d)
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
