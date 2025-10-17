package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

type Field struct {
	Name string
	Type string
	JSON string
}

type StructMeta struct {
	Name   string
	Fields []Field
}

func ParseStructFromFile(filePath, structName string) (*StructMeta, error) {
	fs := token.NewFileSet()
	node, err := parser.ParseFile(fs, filePath, nil, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	var meta StructMeta

	for _, decl := range node.Decls {
		gen, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, spec := range gen.Specs {
			ts, ok := spec.(*ast.TypeSpec)
			if !ok || ts.Name.Name != structName {
				continue
			}
			strct, ok := ts.Type.(*ast.StructType)
			if !ok {
				continue
			}
			meta.Name = structName
			for _, f := range strct.Fields.List {
				if len(f.Names) == 0 {
					continue
				}
				name := f.Names[0].Name
				ftype := exprToString(f.Type)
				jsonName := strings.ToLower(name[:1]) + name[1:]
				meta.Fields = append(meta.Fields, Field{Name: name, Type: ftype, JSON: jsonName})
			}
		}
	}

	return &meta, nil
}

func exprToString(e ast.Expr) string {
	switch v := e.(type) {
	case *ast.Ident:
		return v.Name
	case *ast.SelectorExpr:
		return exprToString(v.X) + "." + v.Sel.Name
	default:
		return ""
	}
}
