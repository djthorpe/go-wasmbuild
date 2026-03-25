package carbon

import (
	"fmt"

	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

// Head returns an <h1>–<h6> styled with the matching Carbon heading token.
// Level 1 maps to cds--heading-06 (largest); level 6 to cds--heading-01.
func Head(level int, args ...any) *text {
	if level < 1 || level > 6 {
		panic(fmt.Sprintf("carbon.Head: level must be 1-6, got %d", level))
	}
	tag := fmt.Sprintf("H%d", level)
	cls := fmt.Sprintf("cds--heading-%02d", 7-level)
	return mvc.NewView(new(text), ViewText, tag, setView, mvc.WithClass(cls), args).(*text)
}
