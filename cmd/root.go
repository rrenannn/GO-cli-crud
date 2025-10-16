package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "crudgen",
	Short: "Gerador automático de CRUD com SQLC + Echo",
	Long:  "CLI para gerar automaticamente repository, service e handler baseados no pacote gerado pelo SQLC.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
}
