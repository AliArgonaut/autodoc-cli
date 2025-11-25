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

	for _, path := range paths {
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

				fmt.Println("Function:", fn.Name.Name)
				fmt.Println("Parameters:", fn.Type.Params)
				fmt.Println("Results:", fn.Type.Results)
			}
			return true
		})
	}
}
