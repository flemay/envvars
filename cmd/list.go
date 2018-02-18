package cmd

import (
	"github.com/flemay/envvars/pkg/envvars"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display the declaration of each environment variable",
	RunE: func(cmd *cobra.Command, args []string) error {
		d, err := envvars.NewDeclaration(declarationFileRootFlag)
		if err != nil {
			return err
		}
		c, err := envvars.List(d, tagsRootFlag...)
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
