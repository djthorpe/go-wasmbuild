package carbon

import (
	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

// HeaderNavGlobal returns the right-aligned global actions container for a header.
func HeaderNavGlobal(args ...any) *navglobal {
	adapted := adaptHeaderGlobalArgs(args...)
	return mvc.NewView(new(navglobal), ViewNavGlobal, "div", setView, append([]any{mvc.WithClass("cds--header__global"), mvc.WithAttr("data-floating-menu-container", "")}, adapted...)...).(*navglobal)
}
