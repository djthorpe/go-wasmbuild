//go:build !(js && wasm)

package js

import "fmt"

///////////////////////////////////////////////////////////////////////////////
// TYPES

type promise struct {
	thenfn    func(value Value) error
	catchfn   func(err error)
	finallyfn func()
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

var (
	errMissingThen = fmt.Errorf("missing 'then' function in promise")
)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// NewPromise creates a new Promise
func NewPromise() *promise {
	return &promise{}
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (p *promise) Try(tryfn func() (Value, error)) {
	if p.thenfn == nil {
		if p.catchfn != nil {
			p.catchfn(errMissingThen)
		} else {
			panic(errMissingThen)
		}
	}
	// Run the function asynchronously, then call the appropriate 'Then',
	// 'Catch' and 'Finally' functions if they are set
	go func() {
		defer func() {
			if p.finallyfn != nil {
				p.finallyfn()
			}
		}()
		if value, err := tryfn(); err != nil {
			if p.catchfn != nil {
				p.catchfn(err)
			}
		} else {
			if p.thenfn != nil {
				if err := p.thenfn(value); err != nil {
					if p.catchfn != nil {
						p.catchfn(err)
					}
				}
			}
		}
	}()
}

// Set the 'then' function which is called after a successful promise
func (p *promise) Then(thenfn func(value Value) error) *promise {
	p.thenfn = thenfn
	return p
}

// Set the 'catch' function which is called after a failed promise
func (p *promise) Catch(catchfn func(err error)) *promise {
	p.catchfn = catchfn
	return p
}

// Set the 'finally' function which is called after the promise is completed
func (p *promise) Finally(finallyfn func()) *promise {
	p.finallyfn = finallyfn
	return p
}
