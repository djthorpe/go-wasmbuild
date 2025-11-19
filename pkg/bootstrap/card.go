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
	BootstrapView
}

type cardgroup struct {
	BootstrapView
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
	c.BootstrapView.View = mvc.NewViewExEx(c, ViewCard, templateCard, args)
	return c
}

func CardGroup(args ...any) *cardgroup {
	c := new(cardgroup)
	c.BootstrapView.View = mvc.NewView(c, ViewCardGroup, "DIV", mvc.WithClass("card-group"), args)
	return c
}

func newCardFromElement(element Element) mvc.View {
	if element.TagName() != "DIV" {
		return nil
	}
	c := new(card)
	c.BootstrapView.View = mvc.NewViewWithElement(c, element)
	return c
}

func newCardGroupFromElement(element Element) mvc.View {
	if element.TagName() != "DIV" {
		return nil
	}
	c := new(cardgroup)
	c.BootstrapView.View = mvc.NewViewWithElement(c, element)
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
		return card.ReplaceSlot("label", mvc.HTML("div"))
	}
	if len(children) > 1 {
		panic("card.Label: only one child element is allowed")
	}
	switch child := children[0].(type) {
	case string:
		return card.ReplaceSlot("label", mvc.HTML("h5", mvc.WithClass("card-title"), mvc.WithInnerText(child)))
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
