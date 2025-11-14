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
	return mvc.NewViewExEx(new(card), ViewCard, templateCard, args).(*card)
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
	card.View = view
}

func (cardgroup *cardgroup) SetView(view mvc.View) {
	cardgroup.View = view
}

func (card *card) Label(children ...any) mvc.View {
	// We only accept one image as a label for the card
	if len(children) == 0 {
		// Clear the content
		return card.View.Label()
	}
	if len(children) > 1 {
		panic("card.Label: only one child element is allowed")
	}
	switch child := children[0].(type) {
	case mvc.View:
		if child.Name() != ViewImage {
			panic(fmt.Sprintf("card.Label: invalid child view type %q", child.Name()))
		}
		card.View.Label(child)
		// TODO: Replace existing label with this child
		fmt.Println("card.Label called", child)
	default:
		panic(fmt.Sprintf("card.Label: invalid child type %T", child))
	}
	return card
}
