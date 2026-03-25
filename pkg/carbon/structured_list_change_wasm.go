//go:build js && wasm

package carbon

import (
	// Packages
	js "github.com/djthorpe/go-wasmbuild/pkg/js"
)

var structuredListChangeBridge = make(map[*structuredList][]js.Func)

func ensureStructuredListChangeBridge(s *structuredList) {
	if s == nil {
		return
	}
	if _, exists := structuredListChangeBridge[s]; exists {
		return
	}
	node, ok := s.Root().JSValue().(js.Value)
	if !ok || node.IsUndefined() || node.IsNull() {
		return
	}
	s.changeBaseline = structuredListActiveRowElement(s)
	listener := js.NewFunc(func(this js.Value, args []js.Value) any {
		current := structuredListActiveRowElement(s)
		same := false
		switch {
		case s.changeBaseline == nil && current == nil:
			same = true
		case s.changeBaseline != nil && current != nil:
			same = current.Equals(s.changeBaseline)
		}
		if same {
			return nil
		}
		s.changeBaseline = current
		init := js.NewObject()
		init.Set("bubbles", true)
		init.Set("composed", true)
		event := js.EventProto.New(EventChange, init)
		target := node
		if current != nil {
			if rowNode, ok := current.JSValue().(js.Value); ok && !rowNode.IsUndefined() && !rowNode.IsNull() {
				target = rowNode
			}
		}
		target.Call("dispatchEvent", event)
		return nil
	})
	structuredListChangeBridge[s] = []js.Func{listener}
	node.Call("addEventListener", EventClick, listener)
	node.Call("addEventListener", "keydown", listener)
}
