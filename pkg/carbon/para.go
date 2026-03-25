package carbon

import (
	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

// Para returns a <p> styled with the Carbon body-01 type token.
func Para(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "P", setView,
		mvc.WithClass("cds--body-01"), args).(*text)
}
