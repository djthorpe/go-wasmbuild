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

// MarkdownLinkResolver rewrites markdown link destinations before they are rendered.
type MarkdownLinkResolver func(string) string

type markdownConfig struct {
	linkResolver MarkdownLinkResolver
}

// MarkdownOpt configures markdown-specific rendering behaviour.
type MarkdownOpt func(*markdownConfig)

var _ mvc.View = (*markdown)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewMarkdown, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(markdown), element, setView)
	})
}

// Markdown creates a block-level markdown view using a DIV root element.
func Markdown(text string, args ...any) mvc.View {
	var cfg markdownConfig
	viewArgs := make([]any, 0, len(args)+2)
	viewArgs = append(viewArgs, mvc.WithClass("markdown"))
	for _, arg := range args {
		switch value := arg.(type) {
		case MarkdownOpt:
			value(&cfg)
		default:
			viewArgs = append(viewArgs, arg)
		}
	}

	children := markdownFallback(text)
	func() {
		defer func() {
			if recover() != nil {
				children = markdownFallback(text)
			}
		}()
		doc := md.Parse(strings.NewReader(text), tokenizer.Pos{})
		children = markdownChildren(doc, cfg)
	}()
	viewArgs = append(viewArgs, children)
	return mvc.NewView(new(markdown), ViewMarkdown, "DIV", setView, viewArgs...)
}

// WithMarkdownLinkResolver applies a link resolver to markdown links.
func WithMarkdownLinkResolver(fn MarkdownLinkResolver) MarkdownOpt {
	return func(cfg *markdownConfig) {
		cfg.linkResolver = fn
	}
}

func markdownFallback(text string) []any {
	trimmed := strings.TrimSpace(text)
	if trimmed == "" {
		return nil
	}
	return []any{Para(trimmed)}
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func markdownChildren(node ast.Node, cfg markdownConfig) []any {
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
		if view := markdownNode(child, cfg); view != nil {
			result = append(result, view)
		}
	}
	flushText()

	return result
}

func markdownNode(node ast.Node, cfg markdownConfig) any {
	if node == nil {
		return nil
	}

	switch n := node.(type) {
	case *md.Document:
		children := markdownChildren(n, cfg)
		return mvc.NewView(new(markdown), ViewMarkdown, "DIV", setView, children)

	case *md.Heading:
		return Head(n.Level(), markdownChildren(n, cfg)...)

	case *md.Paragraph:
		return Para(markdownChildren(n, cfg)...)

	case *md.Strong:
		return Strong(markdownChildren(n, cfg)...)

	case *md.Emphasis:
		return Em(markdownChildren(n, cfg)...)

	case *md.Strikethrough:
		return Deleted(markdownChildren(n, cfg)...)

	case *md.Code:
		return Code(With(ThemeG10), n.Value())

	case *md.CodeBlock:
		codeArgs := []any{With(ThemeG10), n.Content()}
		if language := n.Language(); language != "" {
			codeArgs = append([]any{mvc.WithAttr("data-language", language)}, codeArgs...)
		}
		codeArgs = append(codeArgs, With(CodeWrapText))
		return CodeBlock(codeArgs...)

	case *md.Link:
		href := resolveMarkdownURL(n.URL(), cfg)
		children := markdownChildren(n, cfg)
		return mvc.HTML("A", append([]any{mvc.WithAttr("href", href)}, children...)...)

	case *md.Image:
		return mvc.HTML("IMG", mvc.WithAttr("src", resolveMarkdownURL(n.URL(), cfg)), mvc.WithAttr("alt", n.Alt()))

	case *md.Table:
		return markdownTable(n, cfg)

	case *md.TableHeader:
		return markdownTableHeader(n, cfg)

	case *md.TableRow:
		return markdownTableRow(n, cfg)

	case *md.TableCell:
		children := markdownChildren(n, cfg)
		if len(children) == 1 {
			return children[0]
		}
		return mvc.HTML("SPAN", children...)

	case *md.List:
		children := markdownChildren(n, cfg)
		if n.Ordered() {
			return mvc.HTML("OL", children...)
		}
		return mvc.HTML("UL", children...)

	case *md.ListItem:
		return mvc.HTML("LI", markdownChildren(n, cfg)...)

	case *md.Blockquote:
		return Blockquote(markdownChildren(n, cfg)...)

	case *md.Rule:
		return mvc.HTML("HR", mvc.WithStyle("border:0;border-top:1px solid var(--cds-border-subtle,#c6c6c6);margin:1rem 0"))

	default:
		if len(node.Children()) > 0 {
			return markdownChildren(node, cfg)
		}
		return nil
	}
}

func resolveMarkdownURL(url string, cfg markdownConfig) string {
	if cfg.linkResolver == nil {
		return url
	}
	return cfg.linkResolver(url)
}

func markdownTable(node *md.Table, cfg markdownConfig) mvc.View {
	table := Table()
	if header := node.Header(); header != nil {
		table.Root().QuerySelector("[data-slot=header]").AppendChild(markdownTableHeader(header, cfg).Root())
	}
	body := table.Root().QuerySelector("[data-slot=body]")
	for _, child := range node.Rows() {
		if row, ok := child.(*md.TableRow); ok {
			body.AppendChild(markdownTableRow(row, cfg).Root())
		}
	}
	return Page(mvc.WithStyle("margin-bottom:1rem"), table)
}

func markdownTableHeader(node *md.TableHeader, cfg markdownConfig) mvc.View {
	header := TableHeader()
	args := make([]any, 0, len(node.Children()))
	for _, child := range node.Children() {
		args = append(args, markdownNode(child, cfg))
	}
	header.Content(args...)
	return header
}

func markdownTableRow(node *md.TableRow, cfg markdownConfig) mvc.View {
	row := TableRow()
	args := make([]any, 0, len(node.Children()))
	for _, child := range node.Children() {
		args = append(args, markdownNode(child, cfg))
	}
	row.Content(args...)
	return row
}
