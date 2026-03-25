package carbon

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func init() {
	mvc.RegisterView(ViewDismissibleTag, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(tag), element, setView)
	}, EventTagDismissibleClosed)
}

// DismissibleTag returns a <cds-dismissible-tag> web component.
// An optional leading string sets the text attribute.
func DismissibleTag(args ...any) *tag {
	args = normalizeTagTextArgs(args...)
	normalizeTagArgs(args...)
	return mvc.NewView(new(tag), ViewDismissibleTag, "cds-dismissible-tag", setView, args).(*tag)
}
