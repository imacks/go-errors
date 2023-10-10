package errors

import (
	goerrors "errors"
	"fmt"
	"testing"
)

func TestUnwrapAll(t *testing.T) {
	err := goerrors.New("hello")

	if Unwrap(err) != nil {
		t.Fatalf("expect nil but got %v", Unwrap(err))
	}
	if rooterr := UnwrapAll(err); rooterr != err {
		t.Fatalf("expect root err == err but got %v", rooterr)
	}

	err2 := fmt.Errorf("foo: %w", err)
	if Unwrap(err2) != err {
		t.Fatalf("expect Unwrap(err2) == err but got %v", Unwrap(err2))
	}
	if rooterr := UnwrapAll(err2); rooterr != err {
		t.Fatalf("expect UnwrapAll(err2) == err but got %v", rooterr)
	}

	err3 := fmt.Errorf("bar: %w", err2)
	if Unwrap(err3) != err2 {
		t.Fatalf("expect Unwrap(err3) == err2 but got %v", Unwrap(err3))
	}
	if rooterr := UnwrapAll(err3); rooterr != err {
		t.Fatalf("expect UnwrapAll(err3) == err but got %v", rooterr)
	}
}

func TestUnwrapAll_mixed(t *testing.T) {
	err := goerrors.New("hello")
	err2 := fmt.Errorf("foo: %w", err)
	err3 := &testWrapError{err: err2}

	if Unwrap(err3) != err2 {
		t.Fatalf("expect Unwrap(err3) == err2 but got %v", Unwrap(err3))
	}
	if rooterr := UnwrapAll(err3); rooterr != err {
		t.Fatalf("expect UnwrapAll(err3) == err but got %v", rooterr)
	}
}

type testWrapError struct {
	err error
}

func (w *testWrapError) Error() string { return w.err.Error() }
func (w *testWrapError) Unwrap() error { return w.err }

func TestAs(t *testing.T) {
	refErr := &stringError{msg: "woo"}

	// check we can fish the leaf back
	var mySlot *stringError
	if !As(refErr, &mySlot) {
		t.Fatalf("expect As(...) == true but got false")
	}
	if !goerrors.Is(mySlot, refErr) {
		t.Fatalf("expect As(...) == true but got false")
	}

	// Check we can fish it even if behind something else.
	// Note: this would fail with xerrors.As() because
	// Wrap() uses github.com/pkg/errors which implements
	// Cause() but not Unwrap().
	// This may change with https://github.com/pkg/errors/pull/206.
	wErr := fmt.Errorf("hidden: %w", refErr)
	mySlot = nil
	if !As(wErr, &mySlot) {
		t.Fatalf("expect As(...) == true but got false")
	}
	if !goerrors.Is(mySlot, refErr) {
		t.Fatalf("expect Is(...) == true but got false")
	}

	// Check we can fish the wrapper back.
	refwErr := &stringPrefixError{err: goerrors.New("world"), prefix: "hello"}
	var mywSlot *stringPrefixError
	if !As(refwErr, &mywSlot) {
		t.Fatalf("expect As(...) == true but got false")
	}
	if !goerrors.Is(mywSlot, refwErr) {
		t.Fatalf("expect Is(...) == true but got false")
	}

	// Check that it works even if behind something else.
	wwErr := fmt.Errorf("hidden: %w", refwErr)
	mywSlot = nil
	if !As(wwErr, &mywSlot) {
		t.Fatalf("expect As(...) == true but got false")
	}
	if !goerrors.Is(mywSlot, refwErr) {
		t.Fatalf("expect Is(...) == true but got false")
	}
}

type stringError struct {
	msg string
}

func (m *stringError) Error() string {
	return m.msg
}

type stringPrefixError struct {
	err error
	prefix string
}

func (m *stringPrefixError) Error() string {
	return fmt.Sprintf("%s: %v", m.prefix, m.err)
}
