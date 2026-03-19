//go:build js && wasm

package carbon

import (
	"strings"

	// Packages
	js "github.com/djthorpe/go-wasmbuild/pkg/js"
)

func checkNumberInputValidity(n *numberInput) bool {
	if n == nil {
		return true
	}
	if node, ok := n.Root().JSValue().(js.Value); ok && !node.IsUndefined() && !node.IsNull() {
		return node.Call("checkValidity").Bool()
	}
	if n.Required() && !n.AllowEmpty() && strings.TrimSpace(n.Value()) == "" {
		n.Root().SetAttribute("invalid", "")
		return false
	}
	n.Root().RemoveAttribute("invalid")
	return true
}

func setNumberInputCustomValidity(n *numberInput, message string) {
	if n == nil {
		return
	}
	if node, ok := n.Root().JSValue().(js.Value); ok && !node.IsUndefined() && !node.IsNull() {
		node.Call("setCustomValidity", message)
	}
	message = strings.TrimSpace(message)
	if message == "" {
		n.Root().RemoveAttribute("invalid")
		return
	}
	n.Root().SetAttribute("invalid", "")
	n.SetInvalidText(message)
}
