// package bar is "private" to your project, it cannot be used outside of it
package bar

// this is a method that can be called from within your project
// the example here is really simple, and it would not make much
// sense in 'the real world', but it shows you a bit how visibility works
// if you go to `main.go`, you can only call `bar.SumParam()` and NOT
// `bar.onlyInBar()` since main.go is in a different package
func SumParam(one, two int) int {
	return onlyInBar(one, two)
}

// this method can only be used from within bar
func onlyInBar(one, two int) int {
	return one + two
}

// like sum, but multiply
func MultiplyParam(one, two int) int {
	return one * two
}
