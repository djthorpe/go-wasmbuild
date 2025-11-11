package bootstrap

import (
	"fmt"

	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
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

const (
	ViewRule = "mvc-bs-rule"
)

func init() {
	mvc.RegisterView(ViewRule, newRuleFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func HRule(opt ...mvc.Opt) mvc.View {
	return mvc.NewView(new(rule), ViewRule, "HR", opt...).(*rule)
}

func VRule(opt ...mvc.Opt) mvc.View {
	return mvc.NewView(new(rule), ViewRule, "DIV", append([]mvc.Opt{mvc.WithClass("vr")}, opt...)...)
}

func newRuleFromElement(element Element) mvc.View {
	tagName := element.TagName()
	if tagName != "HR" && tagName != "DIV" {
		panic(fmt.Sprintf("newRuleFromElement: invalid tag name %q", tagName))
	}
	return mvc.NewViewWithElement(new(rule), element)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (rule *rule) SetView(view mvc.View) {
	rule.View = view
}

func (rule *rule) Append(children ...any) mvc.View {
	panic("Append: not supported for rule")
}
