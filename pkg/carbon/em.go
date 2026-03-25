package carbon

import (
	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

// Em returns inline emphasized text.
func Em(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "EM", setView, args).(*text)
}
