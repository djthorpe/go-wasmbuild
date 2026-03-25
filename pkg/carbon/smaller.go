package carbon

import (
	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

// Smaller returns inline smaller supporting text.
func Smaller(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "SMALL", setView, args).(*text)
}
