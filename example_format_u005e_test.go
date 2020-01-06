package font8x8_test

import (
	"github.com/reiver/go-font8x8"

	"fmt"
)

func ExampleType_Format_u005E() {

	fmt.Printf("%z\n", font8x8.U005E)

	// Output:
	//
	// ░░░█░░░░
	// ░░███░░░
	// ░██░██░░
	// ██░░░██░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
}
