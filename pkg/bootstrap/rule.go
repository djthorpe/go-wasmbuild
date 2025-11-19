package bootstrap

import (

	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// text are elements that represent text views
type rule struct {
	BootstrapView
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

func HRule(args ...any) *rule {
	r := new(rule)
	r.BootstrapView.View = mvc.NewView(r, ViewRule, "HR", args...)
	return r
}

func VRule(args ...any) *rule {
	r := new(rule)
	r.BootstrapView.View = mvc.NewView(r, ViewRule, "DIV", mvc.WithClass("vr"), args)
	return r
}

func newRuleFromElement(element Element) mvc.View {
	if element.TagName() != "HR" && element.TagName() != "DIV" {
		return nil
	}
	r := new(rule)
	r.BootstrapView.View = mvc.NewViewWithElement(r, element)
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
