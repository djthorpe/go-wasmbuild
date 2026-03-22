//go:build js && wasm

package carbon

import (
	"strings"

	// Packages
	js "github.com/djthorpe/go-wasmbuild/pkg/js"
)

func checkInputValidity(i *input) bool {
	if i == nil {
		return true
	}
	if node, ok := i.Root().JSValue().(js.Value); ok && !node.IsUndefined() && !node.IsNull() {
		return node.Call("checkValidity").Bool()
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
	if node, ok := i.Root().JSValue().(js.Value); ok && !node.IsUndefined() && !node.IsNull() {
		node.Call("setCustomValidity", message)
	}
	message = strings.TrimSpace(message)
	if message == "" {
		i.Root().RemoveAttribute("invalid")
		return
	}
	i.Root().SetAttribute("invalid", "")
	i.SetInvalidText(message)
}
