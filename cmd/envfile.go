package cmd

import (
	"github.com/flemay/envvars/pkg/envvars"
	"github.com/spf13/cobra"
)

var envfileName string
var overwriteEnvfile bool

var envfileCmd = &cobra.Command{
	Use:   "envfile",
	Short: "Generate an env file based on the declaration file",
	RunE: func(cmd *cobra.Command, args []string) error {
		d, err := envvars.NewDeclaration(declarationFileRootFlag)
		if err != nil {
			return err
		}
		return envvars.Envfile(d, envfileName, overwriteEnvfile, tagsRootFlag...)
	},
}

func init() {
	envfileCmd.Flags().StringVar(&envfileName, "env-file", ".env", "env file to be generated")
	envfileCmd.Flags().BoolVar(&overwriteEnvfile, "overwrite", true, "overwrite the env file if it exists")
	rootCmd.AddCommand(envfileCmd)
}
