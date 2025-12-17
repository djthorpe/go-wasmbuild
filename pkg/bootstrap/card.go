package bootstrap

import (

	// Packages
	"fmt"

	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
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

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

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
	mvc.RegisterView(ViewCard, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(card), element, setView)
	})
	mvc.RegisterView(ViewCardGroup, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(cardgroup), element, setView)
	})
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Card(args ...any) *card {
	return mvc.NewView(new(card), ViewCard, templateCard, setView, args).(*card)
}

func CardGroup(args ...any) *cardgroup {
	return mvc.NewView(new(cardgroup), ViewCardGroup, "DIV", setView, mvc.WithClass("card-group"), args).(*cardgroup)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (card *card) Header(children ...any) *card {
	card.ReplaceSlot("header", mvc.HTML("div", mvc.WithClass("card-header"), children))
	return card
}

func (card *card) Footer(children ...any) *card {
	card.ReplaceSlot("footer", mvc.HTML("div", mvc.WithClass("card-footer"), children))
	return card
}

func (card *card) Content(children ...any) mvc.View {
	card.ReplaceSlot("body", mvc.HTML("div", mvc.WithClass("card-body"), children))
	return card
}

func (card *card) Label(children ...any) mvc.View {
	if len(children) == 0 {
		return card.ReplaceSlot("label", mvc.Placeholder())
	}
	if len(children) > 1 {
		panic("card.Label: only one child element is allowed")
	}
	switch child := children[0].(type) {
	case mvc.View:
		child.Root().ClassList().Add("card-img-top")
		return card.ReplaceSlot("label", child)
	case dom.Element:
		child.ClassList().Add("card-img-top")
		return card.ReplaceSlot("label", child)
	default:
		panic(fmt.Sprintf("card.Label: invalid child type %T", child))
	}
}
