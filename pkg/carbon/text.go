package carbon

import (
	"fmt"
	"strings"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type text struct{ base }
type blockquote struct{ base }

var _ mvc.View = (*text)(nil)
var _ mvc.View = (*blockquote)(nil)
var _ mvc.LabelState = (*blockquote)(nil)

const templateBlockquote = `
	<figure style="margin:0;">
		<blockquote data-slot="body" style="margin:0;padding-inline-start:1rem;border-inline-start:0.25rem solid var(--cds-border-strong);"></blockquote>
		<figcaption data-slot="label" class="cds--helper-text-01" style="margin-block-start:0.5rem;"></figcaption>
	</figure>
`

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewText, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(text), element, setView)
	})
}

// Head returns an <h1>–<h6> styled with the matching Carbon heading token.
// Level 1 maps to cds--heading-06 (largest); level 6 to cds--heading-01.
func Head(level int, args ...any) *text {
	if level < 1 || level > 6 {
		panic(fmt.Sprintf("carbon.Head: level must be 1-6, got %d", level))
	}
	// Carbon heading scale is inverted: h1 → heading-06, h6 → heading-01
	tag := fmt.Sprintf("H%d", level)
	cls := fmt.Sprintf("cds--heading-%02d", 7-level)
	return mvc.NewView(new(text), ViewText, tag, setView, mvc.WithClass(cls), args).(*text)
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

// Blockquote returns a figure containing a styled blockquote and optional label.
func Blockquote(args ...any) *blockquote {
	return mvc.NewView(new(blockquote), ViewText, templateBlockquote, setView, args).(*blockquote)
}

///////////////////////////////////////////////////////////////////////////////
// STATE

// Label returns the citation or attribution for a blockquote.
func (t *blockquote) Label() string {
	if slot := t.Slot("label"); slot != nil {
		return strings.TrimSpace(slot.TextContent())
	}
	return ""
}

// SetLabel sets the citation or attribution for a blockquote.
func (t *blockquote) SetLabel(label string) mvc.View {
	if t.Root().TagName() != "FIGURE" {
		panic("Label can only be applied to Blockquote")
	}
	t.ReplaceSlotChildren("label", label)
	return t
}
