package bootstrap

import (
	// Packages
	"fmt"
	"strconv"

	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type progress struct {
	mvc.View
}

var _ mvc.ViewWithValue = (*progress)(nil)

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
	return mvc.NewView(new(progress), ViewProgress, "DIV", mvc.WithClass("progress"), WithMinMax(0, 100), args).(*progress)
}

func newProgressFromElement(element Element) mvc.View {
	if element.TagName() != "DIV" {
		return nil
	}
	return mvc.NewViewWithElement(new(progress), element)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (progress *progress) SetView(view mvc.View) {
	progress.View = view
}

func (progress *progress) Value() string {
	return progress.Root().GetAttribute("aria-valuenow")
}

func (progress *progress) Min() string {
	return progress.Root().GetAttribute("aria-valuemin")
}

func (progress *progress) Max() string {
	return progress.Root().GetAttribute("aria-valuemax")
}

func (progress *progress) SetValue(value string) mvc.ViewWithValue {
	// Convert value to a floating point number, and ensure it's between min and max
	min, err := strconv.ParseFloat(progress.Min(), 64)
	if err != nil {
		min = 0
	}
	max, err := strconv.ParseFloat(progress.Max(), 64)
	if err != nil {
		max = 100
	}
	val, err := strconv.ParseFloat(value, 64)
	if err != nil {
		val = min
	}
	if val < min {
		val = min
	} else if val > max {
		val = max
	}

	// Set the value
	progress.Root().SetAttribute("aria-valuenow", strconv.FormatFloat(val, 'f', -1, 64))

	// Change the body content width
	percentage := (val - min) / (max - min) * 100
	fmt.Println("Setting progress to ", percentage)

	// Return self
	return progress
}
