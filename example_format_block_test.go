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

	fmt.Print("U+2582 (▂) LOWER ONE QUARTER BLOCK\n\n")
	fmt.Printf("%z\n", font8x8.U2582)

	fmt.Print("U+2583 (▃) LOWER THREE EIGHTHS BLOCK\n\n")
	fmt.Printf("%z\n", font8x8.U2583)

	fmt.Print("U+2584 (▄) LOWER HALF BLOCK\n\n")
	fmt.Printf("%z\n", font8x8.U2584)

	fmt.Print("U+2585 (▅) LOWER FIVE EIGHTHS BLOCK\n\n")
	fmt.Printf("%z\n", font8x8.U2585)

	fmt.Print("U+2586 (▆) LOWER THREE QUARTERS BLOCK\n\n")
	fmt.Printf("%z\n", font8x8.U2586)

	fmt.Print("U+2587 (▇) LOWER SEVEN EIGHTHS BLOCK\n\n")
	fmt.Printf("%z\n", font8x8.U2587)

	fmt.Print("U+2588 (█) FULL BLOCK\n\n")
	fmt.Printf("%z\n", font8x8.U2588)

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
	//
	// U+2582 (▂) LOWER ONE QUARTER BLOCK
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ████████
	// ████████
	//
	// U+2583 (▃) LOWER THREE EIGHTHS BLOCK
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ████████
	// ████████
	// ████████
	//
	// U+2584 (▄) LOWER HALF BLOCK
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ████████
	// ████████
	// ████████
	// ████████
	//
	// U+2585 (▅) LOWER FIVE EIGHTHS BLOCK
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ████████
	// ████████
	// ████████
	// ████████
	// ████████
	//
	// U+2586 (▆) LOWER THREE QUARTERS BLOCK
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ████████
	// ████████
	// ████████
	// ████████
	// ████████
	// ████████
	//
	// U+2587 (▇) LOWER SEVEN EIGHTHS BLOCK
	//
	// ░░░░░░░░
	// ████████
	// ████████
	// ████████
	// ████████
	// ████████
	// ████████
	// ████████
	//
	// U+2588 (█) FULL BLOCK
	//
	// ████████
	// ████████
	// ████████
	// ████████
	// ████████
	// ████████
	// ████████
	// ████████
}

