package content

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	carbon "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	storybook "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/storybook"
)

func CodeView() []any {
	return []any{
		mvc.HTML("DIV", mvc.WithStyle("padding:1.5rem 2rem"), carbon.Head(1, "Source Code")),
		carbon.Section(
			mvc.WithStyle("padding:1.5rem 2rem"),
			carbon.With(carbon.ThemeWhite),
			codeInlineStory(),
			codeSingleStory(),
			codeMultiStory(),
		),
	}
}

func codeInlineStory() dom.Element {
	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem;max-width:56rem;padding:0 1rem"),
		carbon.Para(
			"Run ", carbon.Code("go build ./..."), " to compile all packages, or ",
			carbon.Code("go test ./..."), " to execute the test suite.",
		),
		carbon.Para(
			"Set ", carbon.Code("GOOS=js"), " and ", carbon.Code("GOARCH=wasm"),
			" before building a WebAssembly target.",
		),
	)

	return storybook.Story(
		"Inline",
		"Inline code snippets are used to embed short code fragments within a sentence of prose.",
		canvas,
		nil,
		storybook.Dropdown("Theme", carbon.ThemeWhite, storybook.DefaultThemes, func(theme carbon.Attr) {
			canvas.Apply(carbon.With(theme)...)
		}),
	)
}

func codeSingleStory() dom.Element {
	// Long enough to overflow the container and show the horizontal scroll affordance.
	const cmd = `GOOS=js GOARCH=wasm go build -trimpath -ldflags "-s -w" ` +
		`-o dist/carbon-app/carbon-app.wasm ` +
		`-tags production ` +
		`-gcflags="all=-e" ` +
		`github.com/djthorpe/go-wasmbuild/wasm/carbon-app`

	snippet := carbon.CodeSnippet(cmd)

	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem;max-width:56rem;padding:0 1rem"),
		snippet,
	)

	return storybook.Story(
		"Single Line",
		"Single-line code snippets display a one-line command or expression with a copy button. Long content scrolls horizontally — Carbon does not support text wrapping for this variant.",
		canvas,
		nil,
		storybook.Dropdown("Theme", carbon.ThemeWhite, storybook.DefaultThemes, func(theme carbon.Attr) {
			canvas.Apply(carbon.With(theme)...)
		}),
		storybook.CheckboxGroup("Disabled", "Disable snippet", false, func(v bool) {
			snippet.SetEnabled(!v)
		}),
		storybook.CheckboxGroup("Copy button", "Hide copy button", false, func(v bool) {
			snippet.SetHideCopyButton(v)
		}),
	)
}

func codeMultiStory() dom.Element {
	const src = `package main

import (
	carbon "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc    "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func main() {
	mvc.New(carbon.Section(
		carbon.Header(carbon.HeaderNavGlobal(carbon.Button(carbon.Icon(carbon.IconUserAvatar, carbon.With(carbon.IconSize24))))).SetLabel("#", "My App", ""),
		carbon.SideNav(carbon.SideNavGroup("Content", carbon.SideNavGroupItem("#home", "Home"), carbon.SideNavGroupItem("#about", "About"))),
		carbon.Page(
			carbon.Head(1, "Hello, Carbon"),
			carbon.Para("Built with go-wasmbuild — a Go-native Carbon Design System framework targeting WebAssembly."),
		),
	), carbon.With(carbon.ThemeG90)).Run()
}`

	block := carbon.CodeBlock(src)

	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem;max-width:56rem;padding:0 1rem"),
		block,
	)

	return storybook.Story(
		"Multi Line",
		"Multi-line code blocks display larger code samples and collapse long content behind a \"Show more\" button. Long lines scroll horizontally by default; enable \"Wrap text\" to fold them instead.",
		canvas,
		nil,
		storybook.Dropdown("Theme", carbon.ThemeWhite, storybook.DefaultThemes, func(theme carbon.Attr) {
			canvas.Apply(carbon.With(theme)...)
		}),
		storybook.CheckboxGroup("Disabled", "Disable snippet", false, func(v bool) {
			block.SetEnabled(!v)
		}),
		storybook.CheckboxGroup("Copy button", "Hide copy button", false, func(v bool) {
			block.SetHideCopyButton(v)
		}),
		storybook.CheckboxGroup("Wrap text", "Wrap long lines", false, func(v bool) {
			block.SetWrapText(v)
		}),
	)
}
