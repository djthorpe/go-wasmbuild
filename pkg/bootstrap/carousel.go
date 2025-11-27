package bootstrap

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type carousel struct {
	mvc.View
}

type carouselitem struct {
	mvc.View
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

func init() {
	mvc.RegisterView(ViewCarousel, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(carousel), element, setView)
	})
	mvc.RegisterView(ViewCarouselItem, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(carouselitem), element, setView)
	})
}

const (
	templateCarousel = `
		<div class="carousel slide">
			<div class="carousel-inner" data-slot="body"></div>
			<button class="carousel-control-prev" type="button" data-slot="prev" data-bs-slide="prev">
				<span class="carousel-control-prev-icon" aria-hidden="true"></span>
				<span class="visually-hidden">Previous</span>
			</button>
			<button class="carousel-control-next" type="button" data-slot="next" data-bs-slide="next">
				<span class="carousel-control-next-icon" aria-hidden="true"></span>
				<span class="visually-hidden">Next</span>
			</button>
		</div>
	`
	templateCarouselItem = `
		<div class="carousel-item" data-slot="body">
			<slot name="label"></slot>
		</div>
	`
	templateCarouselItemLabel = `
		<div class="carousel-caption d-none d-md-block" data-slot="body"></div>
	`
)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Carousel(id string, args ...any) *carousel {
	// Create the view
	view := mvc.NewView(new(carousel), ViewCarousel, templateCarousel, setView, mvc.WithAttr("id", id), args).(*carousel)

	// Set prev/next targets
	view.Slot("prev").SetAttribute("data-bs-target", "#"+id)
	view.Slot("next").SetAttribute("data-bs-target", "#"+id)

	// Return self
	return view.Self().(*carousel)
}

func CarouselItem(args ...any) *carouselitem {
	return mvc.NewView(new(carouselitem), ViewCarouselItem, templateCarouselItem, setView, args).(*carouselitem)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (carousel *carousel) Content(args ...any) mvc.View {
	// Convert strings to CarouselItems, make the first one active
	for i, arg := range args {
		switch arg := arg.(type) {
		case string:
			args[i] = CarouselItem(Image(arg))
		}
	}

	// Always make the first item active
	for _, arg := range args {
		if view, ok := arg.(*carouselitem); ok {
			view.Root().ClassList().Add("active")
			break
		}
	}

	return carousel.View.Content(args...)
}

func (carouselitem *carouselitem) Label(args ...any) mvc.View {
	return carouselitem.ReplaceSlot("label", mvc.HTML(templateCarouselItemLabel))
}
