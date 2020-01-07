package font8x8_test

import (
	"github.com/reiver/go-font8x8"

	"fmt"
)

func ExampleType_Format_block() {

	fmt.Print("U+2580 (▀) UPPER HALF BLOCK\n\n")
	fmt.Printf("%z\n", font8x8.U2580)

	fmt.Print("U+2581 (▁) LOWER ONE EIGHTH BLOCK\n\n")
	fmt.Printf("%z\n", font8x8.U2581)

	// Output:
	//
	// U+2580 (▀) UPPER HALF BLOCK
	//
	// ████████
	// ████████
	// ████████
	// ████████
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+2581 (▁) LOWER ONE EIGHTH BLOCK
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ████████
}

