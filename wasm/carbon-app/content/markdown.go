package content

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	carbon "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	storybook "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/storybook"
)

func MarkdownView() []any {
	return []any{
		storybook.PageHeader("Markdown", "Markdown.md"),
		carbon.Section(
			mvc.WithStyle("padding:1.5rem 2rem"),
			carbon.With(carbon.ThemeWhite),
			markdownStory(),
		),
	}
}

func markdownStory() dom.Element {
	const sample = `## This is a markdown heading

Markdown content supports **strong text**, *emphasis*, ~~deletion~~, and inline code like ` + "`const theme = \"white\"`" + `.

- First bullet item
- Second bullet item
- Third bullet item

> Markdown also supports blockquotes for quoted or supporting content.

` + "```go" + `
func Example() string {
    return "hello, markdown"
}
` + "```" + `
`
	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem;max-width:56rem;padding:1rem"),
		carbon.Markdown(sample),
	)

	return storybook.Story(
		"",
		"Markdown content is parsed into Carbon text and semantic HTML elements for rich documentation-style content.",
		canvas,
		nil,
		storybook.Dropdown("Theme", carbon.ThemeWhite, storybook.DefaultThemes, func(theme carbon.Attr) {
			canvas.Apply(carbon.With(theme)...)
		}),
	)
}
