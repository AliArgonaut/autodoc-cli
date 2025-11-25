package services

import (
	"bytes"
	"cli/models"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
)

func GoLanguageService(cwd string, paths []string) {
	fmt.Println(cwd)
	fmt.Println(paths)

	var astFiles []models.ASTFile

	for _, path := range paths {
		files := models.ASTFile{
			Name:       filepath.Base(path),
			Signatures: []models.Signature{},
		}
		fmt.Println("PATH" + path)
		src, err := os.ReadFile(path)
		if err != nil {
			panic(err)
		}

		bn := filepath.Base(path)
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, bn, src, parser.ParseComments)
		if err != nil {
			panic(err)
		}
		ast.Inspect(f, func(n ast.Node) bool {
			if fn, ok := n.(*ast.FuncDecl); ok {
				sig := models.Signature{
					Name:    fn.Name.Name,
					Params:  astFieldListToString(fn.Type.Params),
					Results: astFieldListToString(fn.Type.Results),
				}
				files.Signatures = append(files.Signatures, sig)
			}
			return true
		})
		astFiles = append(astFiles, files)
	}
}

func astFieldListToString(fields *ast.FieldList) []string {
	if fields == nil {
		return nil
	}
	var results []string
	for _, field := range fields.List {
		var buf bytes.Buffer
		printer.Fprint(&buf, token.NewFileSet(), field.Type)
		typeStr := buf.String()

		if len(field.Names) == 0 {
			results = append(results, typeStr)
		} else {
			for _, name := range field.Names {
				results = append(results, name.Name+": "+typeStr)
			}
		}
	}
	return results
}
