package bootstrap

import (
	"fmt"
	"net/url"
	"strings"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// media are elements that represent audio or video views
type media struct {
	mvc.View
}

// mediacontrol is a set of buttons to interact with media elements
type mediacontrol struct {
	mvc.View
}

var _ mvc.View = (*media)(nil)
var _ mvc.View = (*mediacontrol)(nil)

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
	templateMediaControl = `
		<div class="d-flex align-items-center gap-2 my-2 p-1 container-fluid">
			<button type="button" class="btn btn-sm btn-outline-secondary" data-slot="playpause" data-action="playpause">
				<i class="bi bi-play-fill"></i>
			</button>
			<button type="button" class="btn btn-sm btn-outline-secondary" data-slot="stop" data-action="stop">
				<i class="bi bi-stop-fill"></i>
			</button>
			<div class="progress flex-grow-1" role="progressbar" data-slot="progress" style="height: 8px; cursor: pointer;">
				<div class="progress-bar" data-slot="progressbar" style="width: 0%;"></div>
			</div>
			<small class="text-muted font-monospace" data-slot="time" style="min-width: 5em;">-:--:--</small>
		</div>
	`
)

func init() {
	mvc.RegisterView(ViewMedia, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(media), element, setView)
	})
	mvc.RegisterView(ViewMediaControl, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(mediacontrol), element, setView)
	})
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func YouTube(id string, args ...any) *media {
	// TODO: If the id is a full URL, extract the video ID
	// TODO: If the URL is a playlist, handle that too
	view := mvc.NewView(new(media), ViewMedia, templateYouTube, setView, WithColor(Light), mvc.WithAttr(youtubeAttr, "enablejsapi=1"), args)

	// TODO: Add the fullscreen and other attributes to the player and URL

	// Add the src attribute to the iframe, and all the other attributes
	attrs := strings.Fields(view.Root().GetAttribute(youtubeAttr))
	fmt.Printf("Yattrs: %q\n", attrs)
	view.Slot("body").SetAttribute("src", youtubeUrl(id, attrs...))

	// Return the view
	return view.Self().(*media)
}

func Video(src string, args ...any) *media {
	view := mvc.NewView(new(media), ViewMedia, templateVideo, setView, WithColor(Light), mvc.WithAttr("src", src), args)

	view.Slot("body").SetAttribute("src", src)
	view.Slot("body").SetAttribute("controls", "controls")

	// Return the view
	return view.Self().(*media)
}

func MediaControl(args ...any) *mediacontrol {
	view := mvc.NewView(new(mediacontrol), ViewMediaControl, templateMediaControl, setView, args)
	return view.Self().(*mediacontrol)
}

///////////////////////////////////////////////////////////////////////////////
// MEDIA CONTROL METHODS

// SetPlaying updates the play/pause button icon based on playing state
func (mc *mediacontrol) SetPlaying(playing bool) *mediacontrol {
	btn := mc.Slot("playpause")
	if btn == nil {
		return mc
	}
	icon := btn.QuerySelector("i")
	if icon == nil {
		return mc
	}
	if playing {
		icon.SetClassName("bi bi-pause-fill")
	} else {
		icon.SetClassName("bi bi-play-fill")
	}
	return mc
}

// SetProgress updates the progress bar (0.0 to 1.0)
func (mc *mediacontrol) SetProgress(progress float64) *mediacontrol {
	bar := mc.Slot("progressbar")
	if bar == nil {
		return mc
	}
	if progress < 0 {
		progress = 0
	}
	if progress > 1 {
		progress = 1
	}
	bar.SetAttribute("style", fmt.Sprintf("width: %.1f%%;", progress*100))
	return mc
}

// SetTime updates the time display
func (mc *mediacontrol) SetTime(seconds float64) *mediacontrol {
	timeEl := mc.Slot("time")
	if timeEl == nil {
		return mc
	}
	timeEl.SetInnerHTML(formatDuration(seconds))
	return mc
}

// OnPlayPause sets the click handler for the play/pause button
func (mc *mediacontrol) OnPlayPause(handler func(dom.Event)) *mediacontrol {
	if btn := mc.Slot("playpause"); btn != nil {
		btn.AddEventListener("click", handler)
	}
	return mc
}

// OnStop sets the click handler for the stop button
func (mc *mediacontrol) OnStop(handler func(dom.Event)) *mediacontrol {
	if btn := mc.Slot("stop"); btn != nil {
		btn.AddEventListener("click", handler)
	}
	return mc
}

// OnSeek sets the click handler for the progress bar (seeking)
func (mc *mediacontrol) OnSeek(handler func(dom.Event)) *mediacontrol {
	if progress := mc.Slot("progress"); progress != nil {
		progress.AddEventListener("click", handler)
	}
	return mc
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

// formatDuration formats seconds into H:MM:SS or M:SS format
func formatDuration(seconds float64) string {
	totalSeconds := int(seconds)
	hours := totalSeconds / 3600
	minutes := (totalSeconds % 3600) / 60
	secs := totalSeconds % 60

	if hours > 0 {
		return fmt.Sprintf("%d:%02d:%02d", hours, minutes, secs)
	}
	return fmt.Sprintf("%d:%02d", minutes, secs)
}

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
