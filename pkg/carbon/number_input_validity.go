//go:build !(js && wasm)

package carbon

import (
	"strconv"
	"strings"
)

func checkNumberInputValidity(n *numberInput) bool {
	if n == nil {
		return true
	}
	value := strings.TrimSpace(n.Value())
	if value == "" {
		if n.Required() && !n.AllowEmpty() {
			n.Root().SetAttribute("invalid", "")
			return false
		}
		n.Root().RemoveAttribute("invalid")
		return true
	}
	parsed, err := strconv.ParseFloat(value, 64)
	if err != nil {
		n.Root().SetAttribute("invalid", "")
		return false
	}
	if min := strings.TrimSpace(n.Min()); min != "" {
		if bound, err := strconv.ParseFloat(min, 64); err == nil && parsed < bound {
			n.Root().SetAttribute("invalid", "")
			return false
		}
	}
	if max := strings.TrimSpace(n.Max()); max != "" {
		if bound, err := strconv.ParseFloat(max, 64); err == nil && parsed > bound {
			n.Root().SetAttribute("invalid", "")
			return false
		}
	}
	n.Root().RemoveAttribute("invalid")
	return true
}

func setNumberInputCustomValidity(n *numberInput, message string) {
	if n == nil {
		return
	}
	message = strings.TrimSpace(message)
	if message == "" {
		n.Root().RemoveAttribute("invalid")
		return
	}
	n.Root().SetAttribute("invalid", "")
	n.SetInvalidText(message)
}
