package carbon

import (
	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

// Strong returns inline strongly emphasized text.
func Strong(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "STRONG", setView, args).(*text)
}
