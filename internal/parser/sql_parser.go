package parser

import (
	"bufio"
	"os"
	"strings"
)

type QueryMeta struct {
	Name        string
	Type        string // :one, :many, :exec
	HasID       bool
	RequestType string
	ReturnType  string
	Params      string
	Args        string
}

func ParseQueries(filePath, entityName string) ([]QueryMeta, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var queries []QueryMeta
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "-- name:") {
			parts := strings.Fields(line)
			if len(parts) < 3 {
				continue
			}
			q := QueryMeta{
				Name: parts[2],
				Type: "",
			}
			if len(parts) >= 4 {
				q.Type = parts[3]
			}

			// Define padrões básicos para ID / params / return type
			if q.Type == ":one" || q.Type == ":many" {
				q.ReturnType = "db." + entityName
				if q.Type == ":many" {
					q.ReturnType = "[]db." + entityName
				}
			} else if q.Type == ":exec" {
				q.ReturnType = "error"
			}

			// Simples heurística: queries com Update/ByID recebem ID
			if strings.Contains(strings.ToLower(q.Name), "id") || strings.Contains(strings.ToLower(q.Name), "update") {
				q.HasID = true
				if q.Type == ":exec" || q.Type == ":one" {
					q.Params = "id int32"
					q.Args = "int32(id)"
				}
			} else {
				q.Params = "arg db." + entityName + "Params"
				q.Args = "arg"
				q.RequestType = "db." + entityName + "Params"
			}

			queries = append(queries, q)
		}
	}
	return queries, scanner.Err()
}
