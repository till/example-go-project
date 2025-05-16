// this is strictly not necessary, but IMHO tests should not be
// in the same "namespace" as the code, so in go each package
// can also contain a "_test" namespace
package bar_test

import (
	// import from the std lib, there are other libs
	// which have assertions, but you can mostly get
	// away with the stdlib for simple stuff, if you
	// want more: https://github.com/stretchr/testify
	"testing"

	// this is the complete name of the "bar" package
	// it's prefixed by the module name in `go.mod`
	// go makes it pretty easy to understand where code
	// comes from, as this is a path/URL of some kind
	// the github.com/till etc. anbles gopkg docs etc.
	// but you could also call it `module app` instead
	"github.com/till/example-go-project/internal/bar"
)

func TestSum(t *testing.T) {
	if result := bar.SumParam(1, 2); result != 3 {
		t.Errorf("expected '3', got %d", result)
	}
}

// a bit more complex
func TestMultiply(t *testing.T) {
	// this is a for-range loop, over a list of
	// struct we define here in-line
	// it's not very pretty
	for _, testCase := range []struct {
		One    int
		Two    int
		Result int
	}{
		{
			// first test case
			1, 2, 2,
		},
		{
			// second test case
			2, 2, 4,
		},
		{
			// third test case, slightly more verbose
			One:    4,
			Two:    4,
			Result: 16,
		},
	} {
		if result := bar.MultiplyParam(testCase.One, testCase.Two); result != testCase.Result {
			t.Errorf("expected %d, but got: %d", testCase.Result, result)
		}
	}
}
