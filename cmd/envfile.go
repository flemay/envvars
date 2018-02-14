package cmd

import (
	"github.com/flemay/envvars/pkg/envvars"
	"github.com/spf13/cobra"
)

var envfileName string
var overwriteEnvfile bool

var envfileCmd = &cobra.Command{
	Use:   "envfile",
	Short: "Generate an env file based on the definition file",
	RunE: func(cmd *cobra.Command, args []string) error {
		definition, err := envvars.NewDefinition(definitionFileRootFlag)
		if err != nil {
			return err
		}
		return envvars.Envfile(definition, envfileName, overwriteEnvfile)
	},
}

func init() {
	envfileCmd.Flags().StringVar(&envfileName, "env-file", ".env", "env file to be generated")
	envfileCmd.Flags().BoolVar(&overwriteEnvfile, "overwrite", true, "Overwrite the env file if it exists")
}
