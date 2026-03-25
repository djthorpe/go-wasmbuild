package carbon

import (
	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

// Deleted returns inline deleted text.
func Deleted(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "DEL", setView, args).(*text)
}
