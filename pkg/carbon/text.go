package carbon

import (
	"fmt"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type text struct{ base }

const templateBlockquote = `
	<figure style="display:grid;gap:0.5rem;margin:0;">
		<blockquote data-slot="body" style="margin:0;padding:0.75rem 1rem;border-inline-start:0.25rem solid var(--cds-border-strong,#8d8d8d);background:var(--cds-layer-01,#f4f4f4);"></blockquote>
		<figcaption data-slot="label" class="cds--helper-text-01" style="margin:0;padding:0 1rem;color:var(--cds-text-secondary,#525252);"></figcaption>
	</figure>
`

var _ mvc.View = (*text)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewText, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(text), element, setView)
	})
}

// Para returns a <p> styled with the Carbon body-01 type token.
func Para(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "P", setView,
		mvc.WithClass("cds--body-01"), args).(*text)
}

// Compact returns a <p> styled with the Carbon body-compact-01 type token.
func Compact(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "P", setView,
		mvc.WithClass("cds--body-compact-01"), args).(*text)
}

// Lead returns a <p> styled with the Carbon body-02 type token for larger,
// more prominent introductory copy.
func Lead(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "P", setView,
		mvc.WithClass("cds--body-02"), args).(*text)
}

// Deleted returns inline deleted text.
func Deleted(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "DEL", setView, args).(*text)
}

// Highlighted returns inline highlighted text.
func Highlighted(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "MARK", setView, args).(*text)
}

// Smaller returns inline smaller supporting text.
func Smaller(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "SMALL", setView, args).(*text)
}

// Strong returns inline strongly emphasized text.
func Strong(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "STRONG", setView, args).(*text)
}

// Em returns inline emphasized text.
func Em(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "EM", setView, args).(*text)
}

// Code returns inline code text.
func Code(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "CODE", setView, args).(*text)
}

// Blockquote returns a figure containing a styled blockquote and optional label.
func Blockquote(args ...any) *text {
	return mvc.NewView(new(text), ViewText, templateBlockquote, setView, args).(*text)
}

// Head returns an <h1>–<h6> styled with the matching Carbon heading token.
// Level 1 maps to cds--heading-06 (largest); level 6 to cds--heading-01.
func Head(level int, args ...any) *text {
	if level < 1 || level > 6 {
		panic(fmt.Sprintf("carbon.Head: level must be 1–6, got %d", level))
	}
	tag := fmt.Sprintf("H%d", level)
	// Carbon heading scale is inverted: h1 → heading-06, h6 → heading-01
	cls := fmt.Sprintf("cds--heading-%02d", 7-level)
	return mvc.NewView(new(text), ViewText, tag, setView,
		mvc.WithClass(cls), args).(*text)
}

// Label sets the citation or attribution for a blockquote.
func (t *text) Label(children ...any) *text {
	if t.Root().TagName() != "FIGURE" {
		panic("Label can only be applied to Blockquote")
	}
	t.ReplaceSlotChildren("label", children...)
	return t
}
