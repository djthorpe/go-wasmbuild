package main

import (
	"embed"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"runtime"
	"strings"

	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

//go:embed *_examples.go
var embedExamplesFS embed.FS

func sourcecode() string {
	str, err := function(2)
	if err != nil {
		return err.Error()
	}
	return str
}

func function(stack int) (string, error) {
	pc, file, _, ok := runtime.Caller(stack)
	if !ok {
		return "", fmt.Errorf("sourcecode: unable to determine caller")
	}

	data, err := embedExamplesFS.ReadFile(filepath.Base(file))
	if err != nil {
		return "", err
	}

	var function string
	if fn := runtime.FuncForPC(pc); fn != nil {
		function = trimQualifiedName(fn.Name())
	}

	fset := token.NewFileSet()
	fileAST, err := parser.ParseFile(fset, file, string(data), parser.ParseComments)
	if err != nil {
		return "", fmt.Errorf("parse failed: %w", err)
	}

	for _, decl := range fileAST.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Name == nil || fn.Name.Name != function {
			continue
		}
		start := fset.Position(fn.Pos()).Offset
		end := fset.Position(fn.End()).Offset
		if start < 0 || end > len(data) || start >= end {
			break
		}
		return string(data[start:end]), nil
	}

	return "", fmt.Errorf("function %q not found in %s", function, file)
}

func trimQualifiedName(name string) string {
	if i := strings.LastIndex(name, "."); i >= 0 {
		return name[i+1:]
	}
	return name
}

func Example(fn func() (mvc.View, string)) mvc.View {
	view, source := fn()
	source = strings.ReplaceAll(source, ", sourcecode()", "")
	source = strings.ReplaceAll(source, "\t", "  ")
	return bs.Grid(
		view,
		bs.CodeBlock(source, bs.WithColor(bs.Light), bs.WithBorder(), mvc.WithClass("p-3"), mvc.WithStyle("font-size: 0.75em; overflow-x: auto")),
	)
}
