package cmd

import (
	"github.com/flemay/envvars/pkg/envfile"
	"github.com/flemay/envvars/pkg/envvars"
	"github.com/flemay/envvars/pkg/yml"
	"github.com/spf13/cobra"
)

var envfileName string
var overwriteEnvfile bool
var example bool
var removeEnvfile bool

var envfileCmd = &cobra.Command{
	Use:   "envfile",
	Short: "Generate an env file based on the declaration file",
	RunE: func(cmd *cobra.Command, args []string) error {
		d, err := yml.NewDeclaration(declarationFileRootFlag)
		if err != nil {
			return err
		}

		if removeEnvfile {
			return envfile.Remove(envfileName)
		}

		writer := envfile.NewEnvfile(envfileName, example, overwriteEnvfile)
		return envvars.Envfile(d, writer, tagsRootFlag...)
	},
}

func init() {
	envfileCmd.Flags().StringVar(&envfileName, "env-file", ".env", "env file to be generated")
	envfileCmd.Flags().BoolVar(&example, "example", false, "include example values")
	envfileCmd.Flags().BoolVar(&overwriteEnvfile, "overwrite", false, "overwrite the env file if it exists")
	envfileCmd.Flags().BoolVar(&removeEnvfile, "rm", false, "remove the env file")
	rootCmd.AddCommand(envfileCmd)
}
