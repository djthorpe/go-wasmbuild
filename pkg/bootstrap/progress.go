package bootstrap

import (
	// Packages

	"strconv"

	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type progress struct {
	BootstrapView
	striped bool
}

var _ mvc.View = (*progress)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewProgress = "mvc-bs-progress"
)

func init() {
	mvc.RegisterView(ViewProgress, newProgressFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Progress(args ...any) *progress {
	p := new(progress)
	p.BootstrapView.View = mvc.NewView(p, ViewProgress, "DIV", mvc.WithClass("progress"), WithMinMax(0, 100), args)
	return p
}

func StripedProgress(args ...any) *progress {
	view := Progress(args...)
	view.striped = true
	return view
}

func newProgressFromElement(element Element) mvc.View {
	if element.TagName() != "DIV" {
		return nil
	}
	p := new(progress)
	p.BootstrapView.View = mvc.NewViewWithElement(p, element)
	return p
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (progress *progress) Self() mvc.View {
	return progress
}

func (progress *progress) Value() string {
	return progress.Root().GetAttribute("aria-valuenow")
}

func (progress *progress) Set(value string) mvc.View {
	min, max := progress.minMax()
	val, err := strconv.ParseFloat(value, 64)
	if err != nil {
		val = min
	}
	if val < min {
		val = min
	} else if val > max {
		val = max
	}

	// Set the value, and update the view
	progress.Root().SetAttribute("aria-valuenow", strconv.FormatFloat(val, 'f', -1, 64))
	progress.updateView((val - min) / (max - min) * 100)

	// Return self
	return progress
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (progress *progress) minMax() (float64, float64) {
	min, err := strconv.ParseFloat(progress.Root().GetAttribute("aria-valuemin"), 64)
	if err != nil {
		min = 0
	}
	max, err := strconv.ParseFloat(progress.Root().GetAttribute("aria-valuemax"), 64)
	if err != nil {
		max = 100
	}
	if min >= max {
		min, max = max, min
	}

	return min, max
}

func (progress *progress) updateView(pct float64) {
	classes := []string{"progress-bar"}

	// Add animation if striped
	if progress.striped {
		classes = append(classes, "progress-bar-striped", "progress-bar-animated")
	}

	// Propogate the color classes from the root to the progress bar
	prefix := colorPrefixForView(ViewProgress)
	for _, color := range allColors {
		if progress.Root().ClassList().Contains(color.className(prefix)) {
			classes = append(classes, color.className("bg"))
		}
	}

	// Set the content
	progress.Content(
		mvc.HTML("DIV",
			mvc.WithClass(classes...),
			mvc.WithAttr("style", "width: "+strconv.FormatFloat(pct, 'f', 2, 64)+"%;"),
		),
	)
}
