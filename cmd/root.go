package cmd

import (
	"github.com/rrenannn/crudgen/internal/generator"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "crudgen",
}

var generateCmd = &cobra.Command{
	Use:   "generate [entity]",
	Short: "Generate CRUD from SQLC",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		entity := args[0]
		sqlcFile := "db/sqlc/" + entity + ".sql.go"
		sqlFile := "db/query/" + entity + ".sql"
		return generator.Generate(entity, sqlcFile, sqlFile)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
