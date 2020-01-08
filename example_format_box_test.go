package font8x8_test

import (
	"github.com/reiver/go-font8x8"

	"fmt"
)

func ExampleType_Format_box() {

	fmt.Print("U+2500 (─) BOX DRAWINGS LIGHT HORIZONTAL\n\n")
	fmt.Printf("%z\n", font8x8.U2500)

	fmt.Print("U+2501 (━) BOX DRAWINGS HEAVY HORIZONTAL\n\n")
	fmt.Printf("%z\n", font8x8.U2501)

	fmt.Print("U+2502 (│) BOX DRAWINGS LIGHT VERTICAL\n\n")
	fmt.Printf("%z\n", font8x8.U2502)

	fmt.Print("U+2503 (┃) BOX DRAWINGS HEAVY VERTICAL\n\n")
	fmt.Printf("%z\n", font8x8.U2503)

	fmt.Print("U+2504 (┄) BOX DRAWINGS LIGHT TRIPLE DASH HORIZONTAL\n\n")
	fmt.Printf("%z\n", font8x8.U2504)

	fmt.Print("U+2505 (┅) BOX DRAWINGS HEAVY TRIPLE DASH HORIZONTAL\n\n")
	fmt.Printf("%z\n", font8x8.U2505)

	fmt.Print("U+2506 (┆) BOX DRAWINGS LIGHT TRIPLE DASH VERTICAL\n\n")
	fmt.Printf("%z\n", font8x8.U2506)

	fmt.Print("U+2507 (┇) BOX DRAWINGS HEAVY TRIPLE DASH VERTICAL\n\n")
	fmt.Printf("%z\n", font8x8.U2507)

	fmt.Print("U+2508 (┈) BOX DRAWINGS LIGHT QUADRUPLE DASH HORIZONTAL\n\n")
	fmt.Printf("%z\n", font8x8.U2508)

	fmt.Print("U+2509 (┉) BOX DRAWINGS HEAVY QUADRUPLE DASH HORIZONTAL\n\n")
	fmt.Printf("%z\n", font8x8.U2509)

	fmt.Print("U+250A (┊) BOX DRAWINGS LIGHT QUADRUPLE DASH VERTICAL\n\n")
	fmt.Printf("%z\n", font8x8.U250A)

	fmt.Print("U+250B (┋) BOX DRAWINGS HEAVY QUADRUPLE DASH VERTICAL\n\n")
	fmt.Printf("%z\n", font8x8.U250B)

	fmt.Print("U+250C (┌) BOX DRAWINGS LIGHT DOWN AND RIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U250C)

	fmt.Print("U+250D (┍) BOX DRAWINGS DOWN LIGHT AND RIGHT HEAVY\n\n")
	fmt.Printf("%z\n", font8x8.U250D)

	fmt.Print("U+250E (┎) BOX DRAWINGS DOWN HEAVY AND RIGHT LIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U250E)

	fmt.Print("U+250F (┏) BOX DRAWINGS HEAVY DOWN AND RIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U250F)


	// Output:
	//
	// U+2500 (─) BOX DRAWINGS LIGHT HORIZONTAL
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ████████
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+2501 (━) BOX DRAWINGS HEAVY HORIZONTAL
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ████████
	// ████████
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+2502 (│) BOX DRAWINGS LIGHT VERTICAL
	//
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	//
	// U+2503 (┃) BOX DRAWINGS HEAVY VERTICAL
	//
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	//
	// U+2504 (┄) BOX DRAWINGS LIGHT TRIPLE DASH HORIZONTAL
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ██░███░█
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+2505 (┅) BOX DRAWINGS HEAVY TRIPLE DASH HORIZONTAL
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ██░███░█
	// ██░███░█
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+2506 (┆) BOX DRAWINGS LIGHT TRIPLE DASH VERTICAL
	//
	// ░░░█░░░░
	// ░░░░░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░░░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	//
	// U+2507 (┇) BOX DRAWINGS HEAVY TRIPLE DASH VERTICAL
	//
	// ░░░██░░░
	// ░░░░░░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░░░░░░
	// ░░░██░░░
	// ░░░██░░░
	//
	// U+2508 (┈) BOX DRAWINGS LIGHT QUADRUPLE DASH HORIZONTAL
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// █░█░█░█░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+2509 (┉) BOX DRAWINGS HEAVY QUADRUPLE DASH HORIZONTAL
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// █░█░█░█░
	// █░█░█░█░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+250A (┊) BOX DRAWINGS LIGHT QUADRUPLE DASH VERTICAL
	//
	// ░░░░░░░░
	// ░░░█░░░░
	// ░░░░░░░░
	// ░░░█░░░░
	// ░░░░░░░░
	// ░░░█░░░░
	// ░░░░░░░░
	// ░░░█░░░░
	//
	// U+250B (┋) BOX DRAWINGS HEAVY QUADRUPLE DASH VERTICAL
	//
	// ░░░░░░░░
	// ░░░██░░░
	// ░░░░░░░░
	// ░░░██░░░
	// ░░░░░░░░
	// ░░░██░░░
	// ░░░░░░░░
	// ░░░██░░░
	//
	// U+250C (┌) BOX DRAWINGS LIGHT DOWN AND RIGHT
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░█████
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	//
	// U+250D (┍) BOX DRAWINGS DOWN LIGHT AND RIGHT HEAVY
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░█████
	// ░░░█████
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	//
	// U+250E (┎) BOX DRAWINGS DOWN HEAVY AND RIGHT LIGHT
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░█████
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	//
	// U+250F (┏) BOX DRAWINGS HEAVY DOWN AND RIGHT
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░█████
	// ░░░█████
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
}
