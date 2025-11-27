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
		bs.Heading(4, "Responsive Image", mvc.WithClass("mt-4")), Example(Example_Image_001),
		bs.Heading(4, "Thumbnails", mvc.WithClass("mt-4")), Example(Example_Image_002),
		bs.Heading(4, "Rounded", mvc.WithClass("mt-4")), Example(Example_Image_003),
		bs.Heading(4, "With Label (TODO)", mvc.WithClass("mt-4")), Example(Example_Image_004),
		bs.Heading(4, "Carousel (TODO)", mvc.WithClass("mt-4")), Example(Example_Image_005),
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
	return bs.Carousel("image_005"), sourcecode()
}
