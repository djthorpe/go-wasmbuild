package bootstrap

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

// markdown are elements to wrap inline and block markdown content
type markdown struct {
	mvc.View
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

func init() {
	mvc.RegisterView(ViewMarkdown, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(markdown), element, func(self, child mvc.View) {
			self.(*markdown).View = child
		})
	})
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Markdown creates a block-level markdown view using a DIV root element
func Markdown(text string) mvc.View {
	// Parse markdown text
	doc := md.Parse(strings.NewReader(text), tokenizer.Pos{})

	// Convert AST to views
	children := astToViews(doc)

	// Return the markdown block as a DIV
	return mvc.NewView(new(markdown), ViewMarkdown, "DIV", func(self, child mvc.View) {
		self.(*markdown).View = child
	}, mvc.WithClass("markdown"), children)
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

// astToViews converts an AST node and its children to views/elements,
// buffering consecutive text nodes into a single string
func astToViews(node ast.Node) []any {
	if node == nil {
		return nil
	}

	var result []any
	var textBuf strings.Builder

	// Flush buffered text to result
	flushText := func() {
		if textBuf.Len() > 0 {
			result = append(result, textBuf.String())
			textBuf.Reset()
		}
	}

	for _, child := range node.Children() {
		// Check if this is a text node
		if t, ok := child.(*md.Text); ok {
			textBuf.WriteString(t.Value())
			continue
		}
		// Not a text node - flush any buffered text first
		flushText()
		if v := astNodeToView(child); v != nil {
			result = append(result, v)
		}
	}
	// Flush any remaining text
	flushText()

	return result
}

// astNodeToView converts a single AST node to a view or element
func astNodeToView(node ast.Node) any {
	if node == nil {
		return nil
	}

	switch n := node.(type) {
	case *md.Document:
		// Document is a container, return its children
		children := astToViews(n)
		return mvc.NewView(new(markdown), ViewMarkdown, "DIV", func(self, child mvc.View) {
			self.(*markdown).View = child
		}, children)

	case *md.Heading:
		children := astToViews(n)
		return Heading(n.Level(), children)

	case *md.Paragraph:
		children := astToViews(n)
		return Para(children)

	case *md.Strong:
		children := astToViews(n)
		return Strong(children)

	case *md.Emphasis:
		children := astToViews(n)
		return Em(children)

	case *md.Strikethrough:
		children := astToViews(n)
		return Deleted(children)

	case *md.Code:
		return Code(n.Value())

	case *md.CodeBlock:
		return CodeBlock(mvc.WithAttr("data-language", n.Language()), n.Content())

	case *md.Link:
		children := astToViews(n)
		return mvc.NewView(new(text), ViewText, "A", func(self, child mvc.View) {
			self.(*text).View = child
		}, mvc.WithAttr("href", n.URL()), children)

	case *md.Image:
		return mvc.NewView(new(text), ViewText, "IMG", func(self, child mvc.View) {
			self.(*text).View = child
		}, mvc.WithAttr("src", n.URL()), mvc.WithAttr("alt", n.Alt()))

	case *md.List:
		children := astToViews(n)
		if n.Ordered() {
			return mvc.NewView(new(text), ViewText, "OL", func(self, child mvc.View) {
				self.(*text).View = child
			}, children)
		}
		return mvc.NewView(new(text), ViewText, "UL", func(self, child mvc.View) {
			self.(*text).View = child
		}, children)

	case *md.ListItem:
		children := astToViews(n)
		return mvc.NewView(new(text), ViewText, "LI", func(self, child mvc.View) {
			self.(*text).View = child
		}, children)

	case *md.Blockquote:
		children := astToViews(n)
		return Blockquote(children)

	case *md.Rule:
		return mvc.NewView(new(text), ViewText, "HR", func(self, child mvc.View) {
			self.(*text).View = child
		})

	default:
		// Fallback: try to get children if it has any
		if len(node.Children()) > 0 {
			return astToViews(node)
		}
		return nil
	}
}
