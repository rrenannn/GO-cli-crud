package cmd

import "github.com/spf13/cobra"

var moduleName string

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Gera CRUD completo baseados no pacote gerado pelo SQLC.",
	Long:  "Gera CRUD completo baseados no pacote gerado pelo SQLC.",
}

func init() {
	generateCmd.Flags().StringVarP(&moduleName, "module", "m", "", "Nome do módulo")
	generateCmd.MarkFlagRequired("module")
	rootCmd.AddCommand(generateCmd)
}
