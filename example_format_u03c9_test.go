package font8x8_test

import (
	"github.com/reiver/go-font8x8"

	"fmt"
)

func ExampleType_Format_u03C9() {

	fmt.Printf("%z\n", font8x8.U03C9)

	// Output:
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░██░██░░
	// ██░░░██░
	// ██░█░██░
	// ███████░
	// ░██░██░░
	// ░░░░░░░░
}
