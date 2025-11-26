package bootstrap

import (
	// Package imports
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type text struct {
	mvc.View
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	templateBlockquote = `
		<figure>
			<blockquote class="blockquote" data-slot="body"></blockquote>
			<figcaption class="blockquote-footer" data-slot="label"></figcaption>
		</figure>
	`
)

func init() {
	mvc.RegisterView(ViewText, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(text), element, func(self, child mvc.View) {
			self.(*text).View = child
		})
	})
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Para(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "P", func(self, child mvc.View) {
		self.(*text).View = child
	}, args).(*text)
}

func LeadPara(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "P", func(self, child mvc.View) {
		self.(*text).View = child
	}, mvc.WithClass("lead"), args).(*text)
}

func Deleted(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "DEL", func(self, child mvc.View) {
		self.(*text).View = child
	}, args).(*text)
}

func Highlighted(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "MARK", func(self, child mvc.View) {
		self.(*text).View = child
	}, args).(*text)
}

func Smaller(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "SMALL", func(self, child mvc.View) {
		self.(*text).View = child
	}, args).(*text)
}

func Strong(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "STRONG", func(self, child mvc.View) {
		self.(*text).View = child
	}, args).(*text)
}

func Em(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "EM", func(self, child mvc.View) {
		self.(*text).View = child
	}, args).(*text)
}

func Blockquote(args ...any) *text {
	return mvc.NewView(new(text), ViewText, templateBlockquote, func(self, child mvc.View) {
		self.(*text).View = child
	}, args).(*text)
}

func Code(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "CODE", func(self, child mvc.View) {
		self.(*text).View = child
	}, args).(*text)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (text *text) Label(children ...any) mvc.View {
	if text.Root().TagName() != "FIGURE" {
		panic("Label can only be applied to Blockquote")
	}
	return text.ReplaceSlotChildren("label", children...)
}
