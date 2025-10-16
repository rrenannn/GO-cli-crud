package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type TemplateData struct {
	EntityName string
	TableName  string
	ModuleName string
}

func GenerateCRUD(table, moduleName string) error {
	entity := toCamelCase(table)

	files := map[string]string{
		"repository": "templates/repository.tmpl",
		"service":    "templates/service.tmpl",
		"handler":    "templates/handler.tmpl",
		"model":      "templates/model.tmpl",
	}

	for name, path := range files {
		content, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %w", path, err)
		}

		tmpl, err := template.New(name).Parse(string(content))
		if err != nil {
			return fmt.Errorf("erro ao compilar template %s: %w", path, err)
		}

		outputDir := fmt.Sprintf("internal/%s", table)
		os.MkdirAll(outputDir, os.ModePerm)

		outputPath := filepath.Join(outputDir, fmt.Sprintf("%s.go"), name)
		out, err := os.Create(outputPath)
		if err != nil {
			return fmt.Errorf("erro ao criar arquivo %s: %w", outputPath, err)
		}

		defer out.Close()

		if err := tmpl.Execute(out, TemplateData{
			EntityName: entity,
			TableName:  table,
			ModuleName: moduleName,
		}); err != nil {
			return fmt.Errorf("erro ao executar template %s: %w", path, err)
		}
		fmt.Printf("✅ %s gerado em %s\n", strings.Title(name), outputPath)

	}

	fmt.Println("🎉 CRUD completo gerado com sucesso!")
	return nil
}

func toCamelCase(s string) string {
	parts := strings.Split(s, "_")
	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}
	return strings.Join(parts, "")
}
