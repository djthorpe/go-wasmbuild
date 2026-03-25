package carbon

import (
	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

// Lead returns a <p> styled with the Carbon body-02 type token for larger,
// more prominent introductory copy.
func Lead(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "P", setView,
		mvc.WithClass("cds--body-02"), args).(*text)
}
