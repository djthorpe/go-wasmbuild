package js_test

import (
	"errors"
	"testing"

	"github.com/djthorpe/go-wasmbuild/pkg/js"
)

func TestPromise_BasicSuccess(t *testing.T) {
	// Test a simple promise that resolves successfully
	executed := false

	value, err := js.NewPromise(func() (js.Value, error) {
		executed = true
		return js.Undefined(), nil
	}).Wait()

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if !executed {
		t.Error("expected try function to be executed")
	}
	if !value.IsUndefined() {
		t.Error("expected undefined value")
	}
}

func TestPromise_BasicError(t *testing.T) {
	// Test a promise that rejects with an error
	_, err := js.NewPromise(func() (js.Value, error) {
		return js.Undefined(), errors.New("test error")
	}).Wait()

	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestPromise_ThenSuccess(t *testing.T) {
	// Test that Then is called on success
	thenCalled := false

	_, err := js.NewPromise(func() (js.Value, error) {
		return js.Undefined(), nil
	}).Then(func(value js.Value) (js.Value, error) {
		thenCalled = true
		return value, nil
	}).Wait()

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if !thenCalled {
		t.Error("expected Then handler to be called")
	}
}

func TestPromise_ThenNotCalledOnError(t *testing.T) {
	// Test that Then is NOT called when try fails
	thenCalled := false

	js.NewPromise(func() (js.Value, error) {
		return js.Undefined(), errors.New("initial error")
	}).Then(func(value js.Value) (js.Value, error) {
		thenCalled = true
		return value, nil
	}).Wait()

	if thenCalled {
		t.Error("Then handler should not be called when try fails")
	}
}

func TestPromise_ThenErrorTriggersCatch(t *testing.T) {
	// Test that an error in Then triggers Catch
	catchCalled := false

	_, err := js.NewPromise(func() (js.Value, error) {
		return js.Undefined(), nil
	}).Then(func(value js.Value) (js.Value, error) {
		return js.Undefined(), errors.New("then error")
	}).Catch(func(err error) error {
		catchCalled = true
		return nil // recover from error
	}).Wait()

	if err != nil {
		t.Errorf("expected error to be recovered, got %v", err)
	}
	if !catchCalled {
		t.Error("expected Catch handler to be called")
	}
}

func TestPromise_CatchRecovery(t *testing.T) {
	// Test that Catch can recover from an error
	_, err := js.NewPromise(func() (js.Value, error) {
		return js.Undefined(), errors.New("initial error")
	}).Catch(func(err error) error {
		return nil // recover
	}).Wait()

	if err != nil {
		t.Errorf("expected Catch to recover, got %v", err)
	}
}

func TestPromise_CatchPropagation(t *testing.T) {
	// Test that Catch can propagate or transform errors
	_, err := js.NewPromise(func() (js.Value, error) {
		return js.Undefined(), errors.New("original error")
	}).Catch(func(err error) error {
		return errors.New("transformed error")
	}).Wait()

	if err == nil {
		t.Error("expected error to be propagated")
	}
}

func TestPromise_Finally(t *testing.T) {
	// Test that Finally is called on success
	finallyCalled := false

	js.NewPromise(func() (js.Value, error) {
		return js.Undefined(), nil
	}).Finally(func() {
		finallyCalled = true
	}).Wait()

	if !finallyCalled {
		t.Error("expected Finally handler to be called")
	}
}

func TestPromise_FinallyOnError(t *testing.T) {
	// Test that Finally is called even on error
	finallyCalled := false

	js.NewPromise(func() (js.Value, error) {
		return js.Undefined(), errors.New("error")
	}).Finally(func() {
		finallyCalled = true
	}).Wait()

	if !finallyCalled {
		t.Error("expected Finally handler to be called on error")
	}
}

func TestPromise_FullChain(t *testing.T) {
	// Test the full chain: Try -> Then -> Finally
	order := []string{}

	js.NewPromise(func() (js.Value, error) {
		order = append(order, "try")
		return js.Undefined(), nil
	}).Then(func(value js.Value) (js.Value, error) {
		order = append(order, "then")
		return value, nil
	}).Catch(func(err error) error {
		order = append(order, "catch")
		return err
	}).Finally(func() {
		order = append(order, "finally")
	}).Wait()

	expected := []string{"try", "then", "finally"}
	if len(order) != len(expected) {
		t.Errorf("expected %v, got %v", expected, order)
		return
	}
	for i, v := range expected {
		if order[i] != v {
			t.Errorf("expected %v at position %d, got %v", v, i, order[i])
		}
	}
}

func TestPromise_FullChainWithError(t *testing.T) {
	// Test the full chain with error: Try -> Catch -> Finally
	order := []string{}

	js.NewPromise(func() (js.Value, error) {
		order = append(order, "try")
		return js.Undefined(), errors.New("error")
	}).Then(func(value js.Value) (js.Value, error) {
		order = append(order, "then")
		return value, nil
	}).Catch(func(err error) error {
		order = append(order, "catch")
		return nil
	}).Finally(func() {
		order = append(order, "finally")
	}).Wait()

	expected := []string{"try", "catch", "finally"}
	if len(order) != len(expected) {
		t.Errorf("expected %v, got %v", expected, order)
		return
	}
	for i, v := range expected {
		if order[i] != v {
			t.Errorf("expected %v at position %d, got %v", v, i, order[i])
		}
	}
}

func TestPromise_NilTryFunction(t *testing.T) {
	// Test with nil try function - should just call handlers with zero value
	thenCalled := false

	_, err := js.NewPromise(nil).Then(func(value js.Value) (js.Value, error) {
		thenCalled = true
		return value, nil
	}).Wait()

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if !thenCalled {
		t.Error("expected Then to be called even with nil try function")
	}
}

func TestPromise_ChainedWithoutHandlers(t *testing.T) {
	// Test promise with only try function, no handlers
	value, err := js.NewPromise(func() (js.Value, error) {
		return js.Undefined(), nil
	}).Wait()

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if !value.IsUndefined() {
		t.Error("expected undefined value")
	}
}

func TestPromise_ErrorWithoutCatch(t *testing.T) {
	// Test that error is returned when no Catch handler
	_, err := js.NewPromise(func() (js.Value, error) {
		return js.Undefined(), errors.New("unhandled error")
	}).Wait()

	if err == nil {
		t.Error("expected error to be returned")
	}
}

func TestPromise_Done(t *testing.T) {
	// Test that Done callback receives results
	doneCalled := false
	var doneValue js.Value
	var doneErr error

	js.NewPromise(func() (js.Value, error) {
		return js.Undefined(), nil
	}).Done(func(value js.Value, err error) {
		doneCalled = true
		doneValue = value
		doneErr = err
	}).Wait()

	if !doneCalled {
		t.Error("expected Done to be called")
	}
	if doneErr != nil {
		t.Errorf("expected no error in Done, got %v", doneErr)
	}
	if !doneValue.IsUndefined() {
		t.Error("expected undefined value in Done")
	}
}

func TestPromise_Run(t *testing.T) {
	// Test async Run() with Done() callback
	done := make(chan struct{})
	executed := false

	js.NewPromise(func() (js.Value, error) {
		executed = true
		return js.Undefined(), nil
	}).Done(func(value js.Value, err error) {
		close(done)
	}).Run()

	<-done

	if !executed {
		t.Error("expected try function to be executed")
	}
}

func TestPromise_OnceGuard(t *testing.T) {
	// Test that promise only executes once
	executionCount := 0

	p := js.NewPromise(func() (js.Value, error) {
		executionCount++
		return js.Undefined(), nil
	})

	// Call Wait multiple times
	p.Wait()
	p.Wait()
	p.Run()
	p.Wait()

	if executionCount != 1 {
		t.Errorf("expected executor to run once, got %d", executionCount)
	}
}

func TestPromise_FromJSPromise(t *testing.T) {
	// Test FromJSPromise API compatibility
	value, err := js.FromJSPromise(js.Undefined()).Wait()

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if !value.IsUndefined() {
		t.Error("expected undefined value")
	}
}
