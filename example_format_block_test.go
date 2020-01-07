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

	fmt.Print("U+2589 (▉) LEFT SEVEN EIGHTHS BLOCK\n\n")
	fmt.Printf("%z\n", font8x8.U2589)

	fmt.Print("U+258A (▊) LEFT THREE QUARTERS BLOCK\n\n")
	fmt.Printf("%z\n", font8x8.U258A)

	fmt.Print("U+258B (▋) LEFT FIVE EIGHTHS BLOCK\n\n")
	fmt.Printf("%z\n", font8x8.U258B)

	fmt.Print("U+258C (▌) LEFT HALF BLOCK\n\n")
	fmt.Printf("%z\n", font8x8.U258C)

	fmt.Print("U+258D (▍) LEFT THREE EIGHTHS BLOCK\n\n")
	fmt.Printf("%z\n", font8x8.U258D)

	fmt.Print("U+258E (▎) LEFT ONE QUARTER BLOCK\n\n")
	fmt.Printf("%z\n", font8x8.U258E)

	fmt.Print("U+258F (▏) LEFT ONE EIGHTH BLOCK\n\n")
	fmt.Printf("%z\n", font8x8.U258F)

	fmt.Print("U+2590 (▐) RIGHT HALF BLOCK\n\n")
	fmt.Printf("%z\n", font8x8.U2590)

	fmt.Print("U+2591 (░) LIGHT SHADE\n\n")
	fmt.Printf("%z\n", font8x8.U2591)

	fmt.Print("U+2592 (▒) MEDIUM SHADE\n\n")
	fmt.Printf("%z\n", font8x8.U2592)

	fmt.Print("U+2593 (▓) DARK SHADE\n\n")
	fmt.Printf("%z\n", font8x8.U2593)

	fmt.Print("U+2594 (▔) UPPER ONE EIGHTH BLOCK\n\n")
	fmt.Printf("%z\n", font8x8.U2594)

	fmt.Print("U+2595 (▕) RIGHT ONE EIGHTH BLOCK\n\n")
	fmt.Printf("%z\n", font8x8.U2595)

	fmt.Print("U+2596 (▖) QUADRANT LOWER LEFT\n\n")
	fmt.Printf("%z\n", font8x8.U2596)

	fmt.Print("U+2597 (▗) QUADRANT LOWER RIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U2597)

	fmt.Print("U+2598 (▘) QUADRANT UPPER LEFT\n\n")
	fmt.Printf("%z\n", font8x8.U2598)

	fmt.Print("U+2599 (▙) QUADRANT UPPER LEFT AND LOWER LEFT AND LOWER RIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U2599)

	fmt.Print("U+259A (▚) QUADRANT UPPER LEFT AND LOWER RIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U259A)

	fmt.Print("U+259B (▛) QUADRANT UPPER LEFT AND UPPER RIGHT AND LOWER LEFT\n\n")
	fmt.Printf("%z\n", font8x8.U259B)

	fmt.Print("U+259C (▜) QUADRANT UPPER LEFT AND UPPER RIGHT AND LOWER RIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U259C)

	fmt.Print("U+259D (▝) QUADRANT UPPER RIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U259D)

	fmt.Print("U+259E (▞) QUADRANT UPPER RIGHT AND LOWER LEFT\n\n")
	fmt.Printf("%z\n", font8x8.U259E)

	fmt.Print("U+259F (▟) QUADRANT UPPER RIGHT AND LOWER LEFT AND LOWER RIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U259F)

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
	//
	// U+2589 (▉) LEFT SEVEN EIGHTHS BLOCK
	//
	// ███████░
	// ███████░
	// ███████░
	// ███████░
	// ███████░
	// ███████░
	// ███████░
	// ███████░
	//
	// U+258A (▊) LEFT THREE QUARTERS BLOCK
	//
	// ██████░░
	// ██████░░
	// ██████░░
	// ██████░░
	// ██████░░
	// ██████░░
	// ██████░░
	// ██████░░
	//
	// U+258B (▋) LEFT FIVE EIGHTHS BLOCK
	//
	// █████░░░
	// █████░░░
	// █████░░░
	// █████░░░
	// █████░░░
	// █████░░░
	// █████░░░
	// █████░░░
	//
	// U+258C (▌) LEFT HALF BLOCK
	//
	// ████░░░░
	// ████░░░░
	// ████░░░░
	// ████░░░░
	// ████░░░░
	// ████░░░░
	// ████░░░░
	// ████░░░░
	//
	// U+258D (▍) LEFT THREE EIGHTHS BLOCK
	//
	// ███░░░░░
	// ███░░░░░
	// ███░░░░░
	// ███░░░░░
	// ███░░░░░
	// ███░░░░░
	// ███░░░░░
	// ███░░░░░
	//
	// U+258E (▎) LEFT ONE QUARTER BLOCK
	//
	// ██░░░░░░
	// ██░░░░░░
	// ██░░░░░░
	// ██░░░░░░
	// ██░░░░░░
	// ██░░░░░░
	// ██░░░░░░
	// ██░░░░░░
	//
	// U+258F (▏) LEFT ONE EIGHTH BLOCK
	//
	// █░░░░░░░
	// █░░░░░░░
	// █░░░░░░░
	// █░░░░░░░
	// █░░░░░░░
	// █░░░░░░░
	// █░░░░░░░
	// █░░░░░░░
	//
	// U+2590 (▐) RIGHT HALF BLOCK
	//
	// ░░░░████
	// ░░░░████
	// ░░░░████
	// ░░░░████
	// ░░░░████
	// ░░░░████
	// ░░░░████
	// ░░░░████
	//
	// U+2591 (░) LIGHT SHADE
	//
	// █░█░█░█░
	// ░░░░░░░░
	// ░█░█░█░█
	// ░░░░░░░░
	// █░█░█░█░
	// ░░░░░░░░
	// ░█░█░█░█
	// ░░░░░░░░
	//
	// U+2592 (▒) MEDIUM SHADE
	//
	// █░█░█░█░
	// ░█░█░█░█
	// █░█░█░█░
	// ░█░█░█░█
	// █░█░█░█░
	// ░█░█░█░█
	// █░█░█░█░
	// ░█░█░█░█
	//
	// U+2593 (▓) DARK SHADE
	//
	// ████████
	// ░█░█░█░█
	// ████████
	// █░█░█░█░
	// ████████
	// ░█░█░█░█
	// ████████
	// █░█░█░█░
	//
	// U+2594 (▔) UPPER ONE EIGHTH BLOCK
	//
	// ████████
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+2595 (▕) RIGHT ONE EIGHTH BLOCK
	//
	// ░░░░░░░█
	// ░░░░░░░█
	// ░░░░░░░█
	// ░░░░░░░█
	// ░░░░░░░█
	// ░░░░░░░█
	// ░░░░░░░█
	// ░░░░░░░█
	//
	// U+2596 (▖) QUADRANT LOWER LEFT
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ████░░░░
	// ████░░░░
	// ████░░░░
	// ████░░░░
	//
	// U+2597 (▗) QUADRANT LOWER RIGHT
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░████
	// ░░░░████
	// ░░░░████
	// ░░░░████
	//
	// U+2598 (▘) QUADRANT UPPER LEFT
	//
	// ████░░░░
	// ████░░░░
	// ████░░░░
	// ████░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+2599 (▙) QUADRANT UPPER LEFT AND LOWER LEFT AND LOWER RIGHT
	//
	// ████░░░░
	// ████░░░░
	// ████░░░░
	// ████░░░░
	// ████████
	// ████████
	// ████████
	// ████████
	//
	// U+259A (▚) QUADRANT UPPER LEFT AND LOWER RIGHT
	//
	// ████░░░░
	// ████░░░░
	// ████░░░░
	// ████░░░░
	// ░░░░████
	// ░░░░████
	// ░░░░████
	// ░░░░████
	//
	// U+259B (▛) QUADRANT UPPER LEFT AND UPPER RIGHT AND LOWER LEFT
	//
	// ████████
	// ████████
	// ████████
	// ████████
	// ████░░░░
	// ████░░░░
	// ████░░░░
	// ████░░░░
	//
	//
	// U+259C (▜) QUADRANT UPPER LEFT AND UPPER RIGHT AND LOWER RIGHT
	//
	// ████████
	// ████████
	// ████████
	// ████████
	// ░░░░████
	// ░░░░████
	// ░░░░████
	// ░░░░████
	//
	// U+259D (▝) QUADRANT UPPER RIGHT
	//
	// ░░░░████
	// ░░░░████
	// ░░░░████
	// ░░░░████
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+259E (▞) QUADRANT UPPER RIGHT AND LOWER LEFT
	//
	// ░░░░████
	// ░░░░████
	// ░░░░████
	// ░░░░████
	// ████░░░░
	// ████░░░░
	// ████░░░░
	// ████░░░░
	//
	// U+259F (▟) QUADRANT UPPER RIGHT AND LOWER LEFT AND LOWER RIGHT
	//
	// ░░░░████
	// ░░░░████
	// ░░░░████
	// ░░░░████
	// ████████
	// ████████
	// ████████
	// ████████
}
