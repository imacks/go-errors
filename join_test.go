package errors

import (
	"fmt"
	"testing"
)

func TestIsJoin(t *testing.T) {
	if IsJoin(New("blah")) {
		t.Errorf("expect IsJoin() return false")
	}

	if IsJoin(fmt.Errorf("test err: %w", New("blah"))) {
		t.Errorf("expect IsJoin() return false")
	}

	// nil should return false
	if IsJoin(Join()) {
		t.Errorf("expect IsJoin() return false")
	}

	if !IsJoin(Join(New("blah"))) {
		t.Errorf("expect IsJoin() return true")
	}

	if !IsJoin(Join(New("blah"), New("blah2"))) {
		t.Errorf("expect IsJoin() return true")
	}
}

func TestSplit(t *testing.T) {
	if Split(nil) != nil {
		t.Errorf("expect Split() to return nil")
	}

	if Split(New("blah")) != nil {
		t.Errorf("expect Split() to return nil")
	}

	if Split(fmt.Errorf("test err: %w", New("blah"))) != nil {
		t.Errorf("expect Split() to return nil")
	}

	errs := Split(Join(New("blah"), New("blah2")))
	if len(errs) != 2 {
		t.Errorf("expect 2 errors but got %d", len(errs))
	}
	if errs[0].Error() != "blah" || errs[1].Error() != "blah2" {
		t.Errorf("got invalid errors")
	}
}