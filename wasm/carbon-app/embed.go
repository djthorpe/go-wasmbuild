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

	cds "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

//go:embed *_examples.go
var embedExamplesFS embed.FS

// sourcecode returns the source of the function that called sourcecode().
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

	var fname string
	if fn := runtime.FuncForPC(pc); fn != nil {
		fname = trimQualifiedName(fn.Name())
	}

	fset := token.NewFileSet()
	fileAST, err := parser.ParseFile(fset, file, string(data), parser.ParseComments)
	if err != nil {
		return "", fmt.Errorf("parse failed: %w", err)
	}

	for _, decl := range fileAST.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Name == nil || fn.Name.Name != fname {
			continue
		}
		start := fset.Position(fn.Pos()).Offset
		end := fset.Position(fn.End()).Offset
		if start < 0 || end > len(data) || start >= end {
			break
		}
		return string(data[start:end]), nil
	}

	return "", fmt.Errorf("function %q not found in %s", fname, file)
}

func trimQualifiedName(name string) string {
	if i := strings.LastIndex(name, "."); i >= 0 {
		return name[i+1:]
	}
	return name
}

// ExampleRow renders a titled preview+source split: left Col8 shows the live
// view, right Col8 shows the source of the example function.
// An optional description string can be passed as the third argument; it is
// rendered as a compact paragraph between the heading and the grid.
func ExampleRow(title string, fn func() (mvc.View, string), desc ...string) mvc.View {
	view, source := fn()
	source = strings.ReplaceAll(source, ", sourcecode()", "")
	source = strings.ReplaceAll(source, "(mvc.View, string)", "mvc.View")
	source = strings.ReplaceAll(source, "\t", "  ")
	children := []any{
		cds.Heading(5, title),
	}
	if len(desc) > 0 && desc[0] != "" {
		children = append(children, cds.CompactPara(
			mvc.WithAttr("style", "color:var(--cds-text-secondary,#525252);"),
			desc[0],
		))
	}
	children = append(children, cds.Grid(
		mvc.WithAttr("style", "column-gap: var(--cds-spacing-05, 1rem);"),
		cds.Col8(view),
		cds.Col8(cds.CodeMulti(source)),
	))
	return cds.Section(children...)
}
