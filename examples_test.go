package fooer_test

import (
	"fmt"

	"github.com/denis-zakharov/gotest/fooer"
)

func ExampleFooer() {
	foo := fooer.Fooer(3)
	fmt.Println("Expected", foo)
	// Output:
	// Expected Foo
}
