package headings

import (
	dom "github.com/djthorpe/go-wasmbuild"
	"github.com/djthorpe/go-wasmbuild/pkg/carbon"
	"github.com/djthorpe/go-wasmbuild/pkg/mvc"
	storybook "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/storybook"
)

func TextView() []any {
	return []any{
		mvc.HTML("DIV", mvc.WithStyle("padding:1.5rem 2rem"), carbon.Head(1, "Text")),
		carbon.Section(
			mvc.WithStyle("padding:1.5rem 2rem"),
			carbon.With(carbon.ThemeWhite),
			paragraphStory(),
			leadStory(),
			compactStory(),
			blockquoteStory(),
			inlineStylesStory(),
			markdownStory(),
		),
	}
}

func paragraphStory() dom.Element {
	const copy = "Carbon body text is used for supporting copy, descriptions, and longer-form content throughout the interface. It should stay readable across themes and layouts, giving product teams a consistent baseline for explanatory content, inline guidance, and general-purpose narrative text that supports the primary task without competing with headings or controls."
	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem;max-width:48rem;padding:0 1rem"),
		carbon.Para(copy),
	)

	return storybook.Story(
		"Paragraph",
		"Paragraphs use the Carbon body text token for readable supporting content.",
		canvas,
		nil,
		storybook.Dropdown("Theme", carbon.ThemeWhite, storybook.DefaultThemes, func(theme carbon.Attr) {
			canvas.Apply(carbon.With(theme)...)
		}),
	)
}

func leadStory() dom.Element {
	const copy = "Lead text introduces a page, section, or feature with more visual presence than standard paragraph copy. It works well for summaries, opening statements, and the first block of explanatory content when you want to establish context quickly before the reader moves into denser details or supporting information."
	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem;max-width:56rem;padding:0 1rem"),
		carbon.Lead(copy),
	)

	return storybook.Story(
		"Lead",
		"Lead text uses a larger body style for introductory or summary content.",
		canvas,
		nil,
		storybook.Dropdown("Theme", carbon.ThemeWhite, storybook.DefaultThemes, func(theme carbon.Attr) {
			canvas.Apply(carbon.With(theme)...)
		}),
	)
}

func compactStory() dom.Element {
	const copy = "Compact text is useful when space is constrained and the content plays a more supporting role, such as metadata, short descriptions, dense panels, or interface regions where the vertical rhythm needs to stay tight without sacrificing readability."
	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem;max-width:56rem;padding:0 1rem"),
		carbon.Compact(copy),
	)

	return storybook.Story(
		"Compact",
		"Compact text uses a tighter body style for dense supporting content.",
		canvas,
		nil,
		storybook.Dropdown("Theme", carbon.ThemeWhite, storybook.DefaultThemes, func(theme carbon.Attr) {
			canvas.Apply(carbon.With(theme)...)
		}),
	)
}

func blockquoteStory() dom.Element {
	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem;max-width:56rem;padding:0 1rem"),
		carbon.Blockquote(
			"Good content systems need a clear typographic hierarchy so that long-form material remains readable without losing structure or emphasis.",
		).Label("Carbon text storybook"),
	)

	return storybook.Story(
		"Blockquote",
		"Blockquotes highlight quoted or supporting content with an optional citation label.",
		canvas,
		nil,
		storybook.Dropdown("Theme", carbon.ThemeWhite, storybook.DefaultThemes, func(theme carbon.Attr) {
			canvas.Apply(carbon.With(theme)...)
		}),
	)
}

func inlineStylesStory() dom.Element {
	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:0.75rem;max-width:56rem;padding:0 1rem"),
		carbon.Para(carbon.Deleted("This sentence is presented as deleted text.")),
		carbon.Para(carbon.Highlighted("This sentence is highlighted for emphasis.")),
		carbon.Para(carbon.Strong("This sentence uses strong emphasis.")),
		carbon.Para(carbon.Smaller("This sentence uses smaller supporting text.")),
		carbon.Para(carbon.Em("This sentence uses italic emphasis.")),
		carbon.Para(carbon.Code("const theme = \"white\"")),
	)

	return storybook.Story(
		"Inline Styles",
		"Inline text elements cover deletion, highlighting, emphasis, small supporting text, and inline code.",
		canvas,
		nil,
		storybook.Dropdown("Theme", carbon.ThemeWhite, storybook.DefaultThemes, func(theme carbon.Attr) {
			canvas.Apply(carbon.With(theme)...)
		}),
	)
}

func markdownStory() dom.Element {
	const sample = `## Markdown

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
		mvc.WithStyle("display:grid;gap:1rem;max-width:56rem;padding:0 1rem"),
		carbon.Markdown(sample),
	)

	return storybook.Story(
		"Markdown",
		"Markdown content is parsed into Carbon text and semantic HTML elements for rich documentation-style content.",
		canvas,
		nil,
		storybook.Dropdown("Theme", carbon.ThemeWhite, storybook.DefaultThemes, func(theme carbon.Attr) {
			canvas.Apply(carbon.With(theme)...)
		}),
	)
}
