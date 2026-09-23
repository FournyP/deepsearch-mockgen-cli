package services

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

// InterfaceFinder walks a directory tree and parses every non-test Go file to
// collect its interface declarations. Unreadable or unparsable files are skipped.
type InterfaceFinder struct{}

func NewInterfaceFinder() *InterfaceFinder {
	return &InterfaceFinder{}
}

func (f *InterfaceFinder) Find(root string) (map[string]string, error) {
	found := make(map[string]string)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil || !isSourceFile(info.Name()) {
			return nil
		}

		file, err := parser.ParseFile(token.NewFileSet(), path, nil, parser.AllErrors)
		if err != nil {
			return nil
		}

		for _, name := range interfaceNames(file) {
			found[name] = path
		}
		return nil
	})

	return found, err
}

func isSourceFile(name string) bool {
	return strings.HasSuffix(name, ".go") && !strings.HasSuffix(name, "_test.go")
}

func interfaceNames(file *ast.File) []string {
	var names []string
	for _, decl := range file.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			if _, ok := typeSpec.Type.(*ast.InterfaceType); ok {
				names = append(names, typeSpec.Name.Name)
			}
		}
	}
	return names
}
