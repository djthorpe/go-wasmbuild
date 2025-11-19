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

func HRule(args ...any) mvc.View {
	r := new(rule)
	r.View = mvc.NewView(r, ViewRule, "HR", args...)
	return r
}

func VRule(args ...any) mvc.View {
	r := new(rule)
	r.View = mvc.NewView(r, ViewRule, "DIV", mvc.WithClass("vr"), args)
	return r
}

func newRuleFromElement(element Element) mvc.View {
	tagName := element.TagName()
	if tagName != "HR" && tagName != "DIV" {
		panic(fmt.Sprintf("newRuleFromElement: invalid tag name %q", tagName))
	}
	r := new(rule)
	r.View = mvc.NewViewWithElement(r, element)
	return r
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (rule *rule) Self() mvc.View {
	return rule
}

func (rule *rule) Append(children ...any) mvc.View {
	panic("Append: not supported for rule")
}
