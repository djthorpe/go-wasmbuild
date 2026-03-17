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

var _ mvc.View = (*text)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewText, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(text), element, setView)
	})
}

// Para returns a <p> styled with the Carbon body-compact-01 type token.
func Para(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "P", setView,
		mvc.WithClass("cds--body-compact-01"), args).(*text)
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
