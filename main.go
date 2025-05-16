package main

// here are your imports
import (
	"fmt"

	// this is the "bar" package
	"github.com/till/example-go-project/foo"
	"github.com/till/example-go-project/internal/bar"
)

// this is the entry point
func main() {
	fmt.Println("Hello world!")

	// sum two varirables, using the "bar" package
	result := bar.SumParam(1, 2) // should be '3'

	fmt.Printf("result: %d", result)

	// multiply two
	fmt.Printf("result: %d", bar.MultiplyParam(6, 7))

	// call foo package
	fmt.Println(foo.Foo("HELLO WORLD"))
}
