package bs

import (
	"fmt"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
	. "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// text are elements that represent text views
type rule struct {
	View
}

var _ View = (*rule)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewRule = "mvc-bs-rule"
)

func init() {
	RegisterView(ViewRule, newRuleFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Rule(opt ...Opt) *rule {
	return &rule{NewView(ViewRule, "HR", opt...)}
}

func VerticalRule(opt ...Opt) *rule {
	opt = append([]Opt{WithClass("vr")}, opt...)
	return &rule{NewView(ViewRule, "DIV", opt...)}
}

func newRuleFromElement(element Element) View {
	tagName := element.TagName()
	if tagName != "HR" && tagName != "DIV" {
		panic(fmt.Sprintf("newRuleFromElement: invalid tag name %q", tagName))
	}
	return &rule{NewViewWithElement(element)}
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (rule *rule) Append(children ...any) View {
	panic("Append: not supported for rule")
}

func (rule *rule) Content(children ...any) View {
	panic("Content: not supported for rule")
}
