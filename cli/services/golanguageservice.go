package services

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
)

func GoLanguageService(cwd string, paths []string) {
	fmt.Println(cwd)
	fmt.Println(paths)

	src, err := os.ReadFile(paths[0])
	if err != nil {
		panic(err)
	}
	bn := filepath.Base(paths[0])

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, bn, src, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	ast.Inspect(f, func(n ast.Node) bool {
		if fn, ok := n.(*ast.FuncDecl); ok {

			fmt.Printf("Found function: %s\n", fn.Body)
		}
		return true
	})

}
