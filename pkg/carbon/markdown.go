package carbon

import (
	"strings"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	tokenizer "github.com/mutablelogic/go-tokenizer"
	ast "github.com/mutablelogic/go-tokenizer/pkg/ast"
	md "github.com/mutablelogic/go-tokenizer/pkg/markdown"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type markdown struct{ base }

var _ mvc.View = (*markdown)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewMarkdown, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(markdown), element, setView)
	})
}

// Markdown creates a block-level markdown view using a DIV root element.
func Markdown(text string) mvc.View {
	doc := md.Parse(strings.NewReader(text), tokenizer.Pos{})
	children := markdownChildren(doc)
	return mvc.NewView(new(markdown), ViewMarkdown, "DIV", setView, mvc.WithClass("markdown"), children)
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func markdownChildren(node ast.Node) []any {
	if node == nil {
		return nil
	}

	var result []any
	var textBuf strings.Builder

	flushText := func() {
		if textBuf.Len() > 0 {
			result = append(result, textBuf.String())
			textBuf.Reset()
		}
	}

	for _, child := range node.Children() {
		if text, ok := child.(*md.Text); ok {
			textBuf.WriteString(text.Value())
			continue
		}
		flushText()
		if view := markdownNode(child); view != nil {
			result = append(result, view)
		}
	}
	flushText()

	return result
}

func markdownNode(node ast.Node) any {
	if node == nil {
		return nil
	}

	switch n := node.(type) {
	case *md.Document:
		children := markdownChildren(n)
		return mvc.NewView(new(markdown), ViewMarkdown, "DIV", setView, children)

	case *md.Heading:
		return Head(n.Level(), markdownChildren(n)...)

	case *md.Paragraph:
		return Para(markdownChildren(n)...)

	case *md.Strong:
		return Strong(markdownChildren(n)...)

	case *md.Emphasis:
		return Em(markdownChildren(n)...)

	case *md.Strikethrough:
		return Deleted(markdownChildren(n)...)

	case *md.Code:
		return Code(n.Value())

	case *md.CodeBlock:
		codeArgs := []any{n.Content()}
		if language := n.Language(); language != "" {
			codeArgs = append([]any{mvc.WithAttr("data-language", language)}, codeArgs...)
		}
		return CodeBlock(codeArgs...)

	case *md.Link:
		children := markdownChildren(n)
		return mvc.HTML("A", append([]any{mvc.WithAttr("href", n.URL())}, children...)...)

	case *md.Image:
		return mvc.HTML("IMG", mvc.WithAttr("src", n.URL()), mvc.WithAttr("alt", n.Alt()))

	case *md.List:
		children := markdownChildren(n)
		if n.Ordered() {
			return mvc.HTML("OL", children...)
		}
		return mvc.HTML("UL", children...)

	case *md.ListItem:
		return mvc.HTML("LI", markdownChildren(n)...)

	case *md.Blockquote:
		return Blockquote(markdownChildren(n)...)

	case *md.Rule:
		return mvc.HTML("HR", mvc.WithStyle("border:0;border-top:1px solid var(--cds-border-subtle,#c6c6c6);margin:1rem 0"))

	default:
		if len(node.Children()) > 0 {
			return markdownChildren(node)
		}
		return nil
	}
}
