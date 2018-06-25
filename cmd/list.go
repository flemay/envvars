package cmd

import (
	"github.com/flemay/envvars/pkg/envvars"
	"github.com/flemay/envvars/pkg/yml"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display the declaration of each environment variable",
	RunE: func(cmd *cobra.Command, args []string) error {
		reader := yml.NewDeclarationYML(declarationFileRootFlag)
		c, err := envvars.List(reader, tagsRootFlag...)
		if err != nil {
			return err
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Description", "Tags"})
		for _, ev := range c {
			table.Append([]string{ev.Name, ev.Desc, strings.Join(ev.Tags, ", ")})
		}
		table.Render()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
