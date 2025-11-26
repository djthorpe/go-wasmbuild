package bootstrap

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// text are elements that represent text views
type rule struct {
	mvc.View
}

var _ mvc.View = (*rule)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

func init() {
	mvc.RegisterView(ViewRule, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(rule), element, setself)
	})
}

func setself(self, child mvc.View) {
	self.(*rule).View = child
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func HRule(args ...any) *rule {
	return mvc.NewView(new(rule), ViewRule, "HR", setself, args).(*rule)
}

func VRule(args ...any) *rule {
	return mvc.NewView(new(rule), ViewRule, "DIV", setself, mvc.WithClass("vr"), args).(*rule)
}
