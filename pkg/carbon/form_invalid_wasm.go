//go:build js && wasm

package carbon

import (
	// Packages
	js "github.com/djthorpe/go-wasmbuild/pkg/js"
)

var formInvalidBridge = make(map[*form]js.Func)

func ensureFormInvalidBridge(f *form) {
	if f == nil {
		return
	}
	if _, exists := formInvalidBridge[f]; exists {
		return
	}
	node, ok := f.Root().JSValue().(js.Value)
	if !ok || node.IsUndefined() || node.IsNull() {
		return
	}
	listener := js.NewFunc(func(this js.Value, args []js.Value) any {
		target := args[0].Get("target")
		if target.Equal(node) {
			return nil
		}
		init := js.NewObject()
		init.Set("bubbles", false)
		init.Set("composed", false)
		event := js.EventProto.New(EventInvalid, init)
		node.Call("dispatchEvent", event)
		return nil
	})
	formInvalidBridge[f] = listener
	node.Call("addEventListener", EventInvalid, listener, true)
}
