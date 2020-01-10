package font8x8_test

import (
	"github.com/reiver/go-font8x8"

	"fmt"
)

func ExampleType_Format_u03B1() {

	fmt.Printf("%z\n", font8x8.U03B1)

	// Output:
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░███░██░
	// ██░███░░
	// ██░░█░░░
	// ██░███░░
	// ░███░██░
	// ░░░░░░░░
}
