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

// exampleRow is the concrete return type of ExampleRow, allowing ExamplePage
// to distinguish example rows from other content (e.g. LeadPara).
type exampleRow struct{ mvc.View }

// ExamplePage renders a component documentation page with a dark title header
// (ThemeG100) containing the tab bar, and a light body with the panel content.
// Non-ExampleRow rows go into the "Usage" tab; ExampleRow results go into "Demo".
func ExamplePage(title string, rows ...any) mvc.View {
	var usageRows, demoRows []any
	for _, r := range rows {
		if er, ok := r.(*exampleRow); ok {
			demoRows = append(demoRows, er.View)
		} else {
			usageRows = append(usageRows, r)
		}
	}

	const padding = "padding:var(--cds-spacing-07,2rem) var(--cds-spacing-07,2rem) var(--cds-spacing-09,3rem);"
	usageArgs := append([]any{mvc.WithAttr("style", padding)}, usageRows...)
	panes := []any{cds.WithTabsSize("lg"), cds.WithTabsType(cds.TabsContained), cds.TabPane("Usage", cds.Section(usageArgs...))}
	if len(demoRows) > 0 {
		demoArgs := append([]any{mvc.WithAttr("style", padding)}, demoRows...)
		panes = append(panes, cds.TabPane("Demo", cds.Section(demoArgs...)))
	}

	tabBar, panels := cds.TabSetParts(panes...)

	header := cds.Section(
		cds.WithTheme(cds.ThemeG100),
		mvc.WithAttr("style", "padding:var(--cds-spacing-10,4rem) var(--cds-spacing-07,2rem) 0;"),
		cds.Heading(1, title),
		tabBar,
	)
	// Negative margins pull the wrapper flush with the edges of .cds--content
	// which applies padding: var(--cds-spacing-09,3rem) var(--cds-spacing-07,2rem).
	return cds.Section(
		mvc.WithAttr("style", "margin:calc(-1 * var(--cds-spacing-09,3rem)) calc(-1 * var(--cds-spacing-07,2rem)) 0;"),
		header,
		panels,
	)
}

// ExampleRow renders a titled preview+source split: left Col8 shows the live
// view, right Col8 shows the source of the example function.
// An optional description string can be passed as the third argument; it is
// rendered as a compact paragraph between the heading and the grid.
func ExampleRow(title string, fn func() (mvc.View, string), desc ...string) *exampleRow {
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
	return &exampleRow{cds.Section(children...)}
}
