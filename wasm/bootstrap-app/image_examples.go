package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func ImageExamples() mvc.View {
	return bs.Container(
		mvc.WithClass("my-3"),
		Markdown("image_examples.md"),
		bs.Heading(3, "Examples"),
		bs.Heading(4, "Carousel (TODO)", mvc.WithClass("mt-4")), Example(Example_Image_005),
		bs.Heading(4, "Carousel Item Labels", mvc.WithClass("mt-4")), Example(Example_Image_006),
		bs.Heading(4, "Responsive Image", mvc.WithClass("mt-4")), Example(Example_Image_001),
		bs.Heading(4, "Thumbnails", mvc.WithClass("mt-4")), Example(Example_Image_002),
		bs.Heading(4, "Rounded", mvc.WithClass("mt-4")), Example(Example_Image_003),
		bs.Heading(4, "With Label (TODO)", mvc.WithClass("mt-4")), Example(Example_Image_004),
	)
}

func Example_Image_001() (mvc.View, string) {
	// Resize the window to see the image scale to always fit within the window
	return bs.Image("favicon.png"), sourcecode()
}

func Example_Image_002() (mvc.View, string) {
	return bs.Row(
		bs.Col(
			bs.Image("favicon.png", mvc.WithClass("img-thumbnail")),
		),
		bs.Col(
			bs.Image("favicon.png", mvc.WithClass("img-thumbnail")),
		),
		bs.Col(
			bs.Image("favicon.png", mvc.WithClass("img-thumbnail")),
		),
	), sourcecode()
}

func Example_Image_003() (mvc.View, string) {
	return bs.Row(
		bs.Col(
			bs.Image("favicon.png", mvc.WithClass("rounded")),
		),
		bs.Col(
			bs.Image("favicon.png", mvc.WithClass("rounded")),
		),
	), sourcecode()
}

func Example_Image_004() (mvc.View, string) {
	return bs.Image("favicon.png").Label(
		"Favicon Image",
	), sourcecode()
}

func Example_Image_005() (mvc.View, string) {
	return bs.Carousel("image_005",
		"https://cdn.hasselblad.com/f/77891/11656x8742/5e7130d236/b_0494.jpg",
		"https://cdn.hasselblad.com/f/77891/11656x8742/b06b61a912/b_32985.jpg",
		"https://cdn.hasselblad.com/f/77891/11656x8742/00714e9fb0/b_29667.jpg",
	), sourcecode()
}

func Example_Image_006() (mvc.View, string) {
	return bs.Carousel("image_006",
		bs.CarouselItem(
			bs.Image("https://cdn.hasselblad.com/f/77891/11656x8742/5e7130d236/b_0494.jpg"),
		).Label("hello"),
	), sourcecode()
}
