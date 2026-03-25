package carbon

import (
	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

// Highlighted returns inline highlighted text.
func Highlighted(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "MARK", setView, args).(*text)
}
