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
	mvc.ViewWithHeaderFooter
}

type cardgroup struct {
	mvc.View
}

var _ mvc.ViewWithHeaderFooter = (*card)(nil)
var _ mvc.View = (*cardgroup)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewCard      = "mvc-bs-card"
	ViewCardGroup = "mvc-bs-card-group"
)

func init() {
	mvc.RegisterView(ViewCard, newCardFromElement)
	mvc.RegisterView(ViewCardGroup, newCardGroupFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Card(args ...any) *card {
	header := mvc.HTML("DIV", mvc.WithClass("card-header"))
	body := mvc.HTML("DIV", mvc.WithClass("card-body"))
	footer := mvc.HTML("DIV", mvc.WithClass("card-footer"))
	return mvc.NewViewEx(new(card), ViewCard, "DIV", header, body, footer, nil, mvc.WithClass("card"), args).(*card)
}

func CardGroup(args ...any) *cardgroup {
	return mvc.NewView(new(cardgroup), ViewCardGroup, "DIV", mvc.WithClass("card-group"), args).(*cardgroup)
}

func newCardFromElement(element Element) mvc.View {
	tagName := element.TagName()
	if tagName != "DIV" {
		panic(fmt.Sprintf("newCardFromElement: invalid tag name %q", tagName))
	}
	return mvc.NewViewWithElement(new(card), element)
}

func newCardGroupFromElement(element Element) mvc.View {
	tagName := element.TagName()
	if tagName != "DIV" {
		panic(fmt.Sprintf("newCardGroupFromElement: invalid tag name %q", tagName))
	}
	return mvc.NewViewWithElement(new(cardgroup), element)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (card *card) SetView(view mvc.View) {
	card.ViewWithHeaderFooter = view.(mvc.ViewWithHeaderFooter)
}

func (cardgroup *cardgroup) SetView(view mvc.View) {
	cardgroup.View = view
}
