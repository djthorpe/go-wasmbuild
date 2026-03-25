package carbon

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

type formGroup struct{ base }

var _ mvc.View = (*formGroup)(nil)

func init() {
	mvc.RegisterView(ViewFormGroup, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(formGroup), element, setView)
	}, EventFocusBubbled, EventNoFocus)
}

// FormGroup returns a <cds-form-group> web component.
func FormGroup(args ...any) *formGroup {
	return mvc.NewView(new(formGroup), ViewFormGroup, "cds-form-group", setView, args).(*formGroup)
}

func (g *formGroup) AddEventListener(event string, handler func(dom.Event)) mvc.View {
	g.View.AddEventListener(formContainerEvent(event), handler)
	return g
}

func (g *formGroup) RemoveEventListener(event string) mvc.View {
	g.View.RemoveEventListener(formContainerEvent(event))
	return g
}

func (g *formGroup) Label() string {
	return g.Root().GetAttribute("legend-text")
}

func (g *formGroup) SetLabel(text string) *formGroup {
	g.Root().SetAttribute("legend-text", text)
	return g
}
