package carbon

import (
	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

// Compact returns a <p> styled with the Carbon body-compact-01 type token.
func Compact(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "P", setView,
		mvc.WithClass("cds--body-compact-01"), args).(*text)
}
