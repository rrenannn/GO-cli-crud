package generator

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/rrenannn/GO-cli-crud/internal/parser"
)

func Generate(entity, sqlcFile, sqlFile string) error {
	// 1️⃣ Ler struct
	structMeta, err := parser.ParseStructFromFile(sqlcFile, entity)
	if err != nil {
		return err
	}

	// 2️⃣ Ler queries
	queries, err := parser.ParseQueries(sqlFile, entity)
	if err != nil {
		return err
	}

	// 3️⃣ Preparar dados para templates
	data := map[string]interface{}{
		"Entity":         structMeta.Name,
		"EntityLower":    strings.ToLower(structMeta.Name),
		"RequestFields":  filterRequestFields(structMeta.Fields),
		"ResponseFields": structMeta.Fields,
		"Queries":        queries,
	}

	// 4️⃣ Criar diretório
	outputDir := filepath.Join("internal", strings.ToLower(entity))
	os.MkdirAll(outputDir, 0755)

	// 5️⃣ Gerar arquivos
	files := []string{"model", "repository", "service", "handler"}
	for _, f := range files {
		tmpl, err := template.ParseFiles(filepath.Join("internal/generator/templates", f+".tmpl"))
		if err != nil {
			return err
		}
		out, err := os.Create(filepath.Join(outputDir, f+".go"))
		if err != nil {
			return err
		}
		defer out.Close()
		tmpl.Execute(out, data)
	}

	return nil
}

func filterRequestFields(fields []parser.Field) []parser.Field {
	var req []parser.Field
	for _, f := range fields {
		if f.Name != "ID" && f.Name != "CreatedAt" && f.Name != "UpdatedAt" {
			req = append(req, f)
		}
	}
	return req
}
