package bootstrap

///////////////////////////////////////////////////////////////////////////////
// TYPES

// Color defines the color for components and backgrounds
type Color string

///////////////////////////////////////////////////////////////////////////////
// CONSTANTS

const (
	TRANSPARENT      Color = ""
	PRIMARY          Color = "primary"
	PRIMARY_SUBTLE   Color = "primary-subtle"
	SECONDARY        Color = "secondary"
	SECONDARY_SUBTLE Color = "secondary-subtle"
	SUCCESS          Color = "success"
	SUCCESS_SUBTLE   Color = "success-subtle"
	DANGER           Color = "danger"
	DANGER_SUBTLE    Color = "danger-subtle"
	WARNING          Color = "warning"
	WARNING_SUBTLE   Color = "warning-subtle"
	INFO             Color = "info"
	INFO_SUBTLE      Color = "info-subtle"
	LIGHT            Color = "light"
	LIGHT_SUBTLE     Color = "light-subtle"
	DARK             Color = "dark"
	DARK_SUBTLE      Color = "dark-subtle"
	WHITE            Color = "white"
	BLACK            Color = "black"
)

///////////////////////////////////////////////////////////////////////////////
// PRIVATE FUNCTIONS

func (color Color) className(prefix string) string {
	return prefix + "-" + string(color)
}
