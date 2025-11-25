package utils

import (
	"cli/models"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
)

func GetGoASTFiles(cwd string, paths []string) []models.ASTFile {

	var astFiles []models.ASTFile

	for _, path := range paths {
		files := models.ASTFile{
			Name:       filepath.Base(path),
			Signatures: []models.Signature{},
		}
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
					Params:  ASTFieldListToString(fn.Type.Params),
					Results: ASTFieldListToString(fn.Type.Results),
				}
				files.Signatures = append(files.Signatures, sig)
			}
			return true
		})
		astFiles = append(astFiles, files)
	}
	return astFiles
}
