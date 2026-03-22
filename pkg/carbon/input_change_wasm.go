//go:build js && wasm

package carbon

import (
	// Packages
	js "github.com/djthorpe/go-wasmbuild/pkg/js"
)

var inputChangeBridge = make(map[*input]js.Func)

func ensureInputChangeBridge(i *input) {
	if i == nil {
		return
	}
	if _, exists := inputChangeBridge[i]; exists {
		return
	}
	node, ok := i.Root().JSValue().(js.Value)
	if !ok || node.IsUndefined() || node.IsNull() {
		return
	}
	listener := js.NewFunc(func(this js.Value, args []js.Value) any {
		current := i.Value()
		if current == i.changeBaseline {
			return nil
		}
		i.changeBaseline = current
		init := js.NewObject()
		init.Set("bubbles", true)
		init.Set("composed", true)
		event := js.EventProto.New(EventChange, init)
		node.Call("dispatchEvent", event)
		return nil
	})
	inputChangeBridge[i] = listener
	node.Call("addEventListener", EventNoFocus, listener)
}
