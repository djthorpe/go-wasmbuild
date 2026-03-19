//go:build !(js && wasm)

package carbon

import "strings"

func checkInputValidity(i *input) bool {
	if i == nil {
		return true
	}
	if i.Required() && strings.TrimSpace(i.Value()) == "" {
		i.Root().SetAttribute("invalid", "")
		return false
	}
	i.Root().RemoveAttribute("invalid")
	return true
}

func setInputCustomValidity(i *input, message string) {
	if i == nil {
		return
	}
	message = strings.TrimSpace(message)
	if message == "" {
		i.Root().RemoveAttribute("invalid")
		return
	}
	i.Root().SetAttribute("invalid", "")
	i.SetInvalidText(message)
}
