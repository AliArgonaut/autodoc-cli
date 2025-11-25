package utils

import (
	"bytes"
	"go/ast"
	//"go/printer"
	//"go/token"
)

func ASTFieldListToString(fields *ast.FieldList) []string {
	if fields == nil {
		return nil
	}
	var results []string
	for _, field := range fields.List {
		var buf bytes.Buffer
		//printer.Fprint(&buf, token.NewFileSet(), field.Type)
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
