package carbon

import (
	"strings"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type blockquote struct{ base }

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
	mvc.RegisterView(ViewBlockquote, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(blockquote), element, setView)
	})
}

// Blockquote returns a figure containing a styled blockquote and optional label.
func Blockquote(args ...any) *blockquote {
	return mvc.NewView(new(blockquote), ViewBlockquote, templateBlockquote, setView, args).(*blockquote)
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