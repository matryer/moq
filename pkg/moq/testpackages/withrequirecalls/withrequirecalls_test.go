package withrequirecalls

import (
	"strings"
	"testing"
)

func TestRequireCalls(t *testing.T) {
	r := RequireCallsMock{}

	// no funcs configured, expect no error
	err := r.RequireCalls()
	if err != nil {
		t.Errorf("All methods are nil, but got error: %s", err)
	}

	// configure both funcs
	r.GetExampleFunc = func() {}
	r.SetExampleFunc = func() {}

	// call only one of them
	r.GetExample()

	// expect error
	err = r.RequireCalls()
	if err == nil {
		t.Fatalf("Expected error, since not all methods have been called")
	}

	// expect error to contain SetExample
	if expected := "RequireCallsMock.SetExample is non-nil but was never called"; !strings.Contains(err.Error(), expected) {
		t.Errorf("Unexpected error message:\n\tExpected: %s\n\tGot:      %s", expected, err)
	}

	// call SetExample
	r.SetExample()

	// all configured funcs have been called, no error
	err = r.RequireCalls()
	if err != nil {
		t.Errorf("All methods should have been called, but got error: %s", err)
	}
}
