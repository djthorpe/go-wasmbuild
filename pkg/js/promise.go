//go:build !(js && wasm)

package js

///////////////////////////////////////////////////////////////////////////////
// TYPES

type promise struct {
	thenfn    func(value Value) any
	catchfn   func(reason Value) any
	finallyfn func()
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// NewPromise creates a new Promise
func NewPromise() *promise {
	return &promise{}
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (p *promise) Then(thenfn func(value Value) any) *promise {
	p.thenfn = thenfn
	return p
}

func (p *promise) Catch(catchfn func(reason Value) any) *promise {
	p.catchfn = catchfn
	return p
}

func (p *promise) Finally(finallyfn func()) *promise {
	p.finallyfn = finallyfn
	return p
}
