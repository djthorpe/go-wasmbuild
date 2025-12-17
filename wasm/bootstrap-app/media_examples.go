package main

import (
	// Packages

	"fmt"

	dom "github.com/djthorpe/go-wasmbuild"
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func MediaExamples() mvc.View {
	return bs.Container(
		mvc.WithClass("my-3"),
		Markdown("media_examples.md"),
		ExampleCard("Media Controls", Example_Media_004),
		ExampleCard("YouTube Video Embedding", Example_Media_001),
		ExampleCard("Without Controls", Example_Media_002),
		ExampleCard("Native Video", Example_Media_003),
	)
}

func Example_Media_001() (mvc.View, string) {
	return bs.YouTube("cbB3QEwWMlA"), sourcecode()
}

func Example_Media_002() (mvc.View, string) {
	return bs.YouTube("cbB3QEwWMlA", bs.WithoutControls()), sourcecode()
}

func Example_Media_003() (mvc.View, string) {
	return bs.Video("http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4"), sourcecode()
}

func Example_Media_004() (mvc.View, string) {
	return bs.MediaControl(
		bs.WithColor(bs.Dark), bs.WithTheme(bs.Dark), bs.WithBorder(),
	).AddEventListener(bs.EventMediaPlayPause, func(evt dom.Event) {
		view := mvc.ViewFromEvent(evt, bs.ViewMediaControl)
		fmt.Println("Media Play/Pause event from view:", view.ID())
	}), sourcecode()
}
