package utils

import (
	"github.com/mazezen/echoframe/cmd/auto/dt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
)

func ParseModels(modelsPath string) ([]dt.ModelInfo, error) {
	var models []dt.ModelInfo
	err := filepath.Walk(modelsPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".go" {
			fset := token.NewFileSet()
			node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
			if err != nil {
				return err
			}
			for _, decl := range node.Decls {
				gen, ok := decl.(*ast.GenDecl)
				if !ok || gen.Tok != token.TYPE {
					continue
				}
				for _, spec := range gen.Specs {
					typeSpec, ok := spec.(*ast.TypeSpec)
					if !ok {
						continue
					}
					structType, ok := typeSpec.Type.(*ast.StructType)
					if !ok {
						continue
					}
					var fields []string
					for _, field := range structType.Fields.List {
						fields = append(fields, field.Names[0].Name)
					}
					models = append(models, dt.ModelInfo{Name: typeSpec.Name.Name, Fields: fields})
				}
			}
		}
		return nil
	})
	return models, err
}
