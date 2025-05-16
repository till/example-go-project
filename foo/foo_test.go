package foo_test

import (
	"testing"

	"github.com/till/example-go-project/foo"
)

// a type — only for the "foo_test" package
// it's lower-case, remember?
// you probably don't want to use testData
// outside of your test suite
type testData struct {
	Input  string
	Output string
}

func TestFoo(t *testing.T) {
	// here's another test method, this time a little nicer setup than in
	// `bar_test` tests, — a bit more organized and nicer to read maybe

	for _, tc := range dataProvider() {
		if result := foo.Foo(tc.Input); result != tc.Output {
			t.Errorf("failed — expected: %s, got: %s", tc.Output, result)
		}
	}

}

// returns the test cases used to ... "test"
func dataProvider() []testData {
	return []testData{
		{
			Input:  "hello",
			Output: "hello",
		},
		{
			Input:  "world",
			Output: "world",
		},
	}
}
