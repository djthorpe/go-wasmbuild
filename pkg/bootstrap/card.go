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

type card struct {
	mvc.View
}

type cardgroup struct {
	mvc.View
}

var _ mvc.View = (*cardgroup)(nil)
var _ mvc.View = (*card)(nil)
var _ mvc.View = (*card)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewCard      = "mvc-bs-card"
	ViewCardGroup = "mvc-bs-cardgroup"
)

const (
	templateCard = `
		<div class="card">
			<slot name="header"><!-- Header --></slot>
			<slot name="label"><!-- Image --></slot>
			<slot><!-- Body --></slot>
			<slot name="footer"><!-- Footer --></slot>
		</div>
	`
)

func init() {
	mvc.RegisterView(ViewCard, newCardFromElement)
	mvc.RegisterView(ViewCardGroup, newCardGroupFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Card(args ...any) *card {
	c := new(card)
	c.View = mvc.NewViewExEx(c, ViewCard, templateCard, args)
	return c
}

func CardGroup(args ...any) *cardgroup {
	c := new(cardgroup)
	c.View = mvc.NewView(c, ViewCardGroup, "DIV", mvc.WithClass("card-group"), args)
	return c
}

func newCardFromElement(element Element) mvc.View {
	tagName := element.TagName()
	if tagName != "DIV" {
		panic(fmt.Sprintf("newCardFromElement: invalid tag name %q", tagName))
	}
	c := new(card)
	c.View = mvc.NewViewWithElement(c, element)
	return c
}

func newCardGroupFromElement(element Element) mvc.View {
	tagName := element.TagName()
	if tagName != "DIV" {
		panic(fmt.Sprintf("newCardGroupFromElement: invalid tag name %q", tagName))
	}
	c := new(cardgroup)
	c.View = mvc.NewViewWithElement(c, element)
	return c
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (card *card) Self() mvc.View {
	return card
}

func (cardgroup *cardgroup) Self() mvc.View {
	return cardgroup
}

func (card *card) Header(children ...any) mvc.View {
	return card.ReplaceSlot("header", mvc.HTML("div", mvc.WithClass("card-header"), children))
}

func (card *card) Footer(children ...any) mvc.View {
	return card.ReplaceSlot("footer", mvc.HTML("div", mvc.WithClass("card-footer"), children))
}

func (card *card) Content(children ...any) mvc.View {
	return card.ReplaceSlot("", mvc.HTML("div", mvc.WithClass("card-body"), children))
}

func (card *card) Label(children ...any) mvc.View {
	if len(children) == 0 {
		return card.View.Label()
	}
	if len(children) > 1 {
		panic("card.Label: only one child element is allowed")
	}
	switch child := children[0].(type) {
	case mvc.View:
		child.Root().ClassList().Add("card-img-top")
		return card.ReplaceSlot("label", child)
	case Element:
		child.ClassList().Add("card-img-top")
		return card.ReplaceSlot("label", child)
	default:
		panic(fmt.Sprintf("card.Label: invalid child type %T", child))
	}
}
