package bootstrap

import (
	// Packages

	"fmt"
	"net/url"
	"strings"

	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// media are elements that represent audio or video views
type media struct {
	mvc.View
}

var _ mvc.View = (*media)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	youtubeEmbedURL = "https://www.youtube.com/embed/"
	youtubeAttr     = "data-yt"
	templateYouTube = `
		<div class="ratio ratio-16x9">
		  <iframe data-slot="body" frameborder="0" allowfullscreen></iframe>
		</div>
	`
	templateVideo = `
		<div class="ratio ratio-16x9">
		  <video data-slot="body" frameborder="0" allowfullscreen></video>
		</div>
	`
)

func init() {
	mvc.RegisterView(ViewMedia, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(media), element, func(self, child mvc.View) {
			self.(*media).View = child
		})
	})
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func YouTube(id string, args ...any) *media {
	// TODO: If the id is a full URL, extract the video ID
	// TODO: If the URL is a playlist, handle that too
	view := mvc.NewView(new(media), ViewMedia, templateYouTube, func(self, child mvc.View) {
		self.(*media).View = child
	}, WithColor(Light), mvc.WithAttr(youtubeAttr, "enablejsapi=1"), args)

	// TODO: Add the fullscreen and other attributes to the player and URL

	// Add the src attribute to the iframe, and all the other attributes
	attrs := strings.Fields(view.Root().GetAttribute(youtubeAttr))
	fmt.Printf("Yattrs: %q\n", attrs)
	view.Slot("body").SetAttribute("src", youtubeUrl(id, attrs...))

	// Return the view
	return view.Self().(*media)
}

func Video(src string, args ...any) *media {
	view := mvc.NewView(new(media), ViewMedia, templateVideo, func(self, child mvc.View) {
		self.(*media).View = child
	}, WithColor(Light), mvc.WithAttr("src", src), args)

	view.Slot("body").SetAttribute("src", src)
	view.Slot("body").SetAttribute("controls", "controls")

	// Return the view
	return view.Self().(*media)
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func youtubeUrl(id string, attrs ...string) string {
	url, err := url.Parse(youtubeEmbedURL + id)
	if err != nil {
		panic("YouTube: " + err.Error())
	}

	// Handle attributes
	q := url.Query()
	for _, attr := range attrs {
		parts := strings.SplitN(attr, "=", 2)
		if len(parts) == 2 {
			q.Set(parts[0], parts[1])
		} else {
			q.Set(parts[0], "1")
		}
	}
	url.RawQuery = q.Encode()

	// Return the URL string
	return url.String()
}

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

func WithAutoplay() mvc.Opt {
	return func(opts mvc.OptSet) error {
		// Append autoplay=1 to existing attribute
		attr := append(strings.Fields(opts.Attr(youtubeAttr)), "autoplay=1")

		// Return updated attribute
		return mvc.WithAttr(youtubeAttr, strings.Join(attr, " "))(opts)
	}
}

func WithoutControls() mvc.Opt {
	return func(opts mvc.OptSet) error {
		// Append autoplay=1 to existing attribute
		attr := append(strings.Fields(opts.Attr(youtubeAttr)), "controls=0")

		// Return updated attribute
		return mvc.WithAttr(youtubeAttr, strings.Join(attr, " "))(opts)
	}
}

func WithoutKeyboardControls() mvc.Opt {
	return func(opts mvc.OptSet) error {
		// Append autoplay=1 to existing attribute
		attr := append(strings.Fields(opts.Attr(youtubeAttr)), "disablekb=1")

		// Return updated attribute
		return mvc.WithAttr(youtubeAttr, strings.Join(attr, " "))(opts)
	}
}

func WithoutFullscreen() mvc.Opt {
	return func(opts mvc.OptSet) error {
		// Append autoplay=1 to existing attribute
		attr := append(strings.Fields(opts.Attr(youtubeAttr)), "fs=0")

		// Return updated attribute
		return mvc.WithAttr(youtubeAttr, strings.Join(attr, " "))(opts)
	}
}
