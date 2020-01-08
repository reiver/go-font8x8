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

	fmt.Print("U+2510 (┐) BOX DRAWINGS LIGHT DOWN AND LEFT\n\n")
	fmt.Printf("%z\n", font8x8.U2510)

	fmt.Print("U+2511 (┑) BOX DRAWINGS DOWN LIGHT AND LEFT HEAVY\n\n")
	fmt.Printf("%z\n", font8x8.U2511)

	fmt.Print("U+2512 (┒) BOX DRAWINGS DOWN HEAVY AND LEFT LIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U2512)

	fmt.Print("U+2513 (┓) BOX DRAWINGS HEAVY DOWN AND LEFT\n\n")
	fmt.Printf("%z\n", font8x8.U2513)

	fmt.Print("U+2514 (└) BOX DRAWINGS LIGHT UP AND RIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U2514)

	fmt.Print("U+2515 (┕) BOX DRAWINGS UP LIGHT AND RIGHT HEAVY\n\n")
	fmt.Printf("%z\n", font8x8.U2515)

	fmt.Print("U+2516 (┖) BOX DRAWINGS UP HEAVY AND RIGHT LIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U2516)

	fmt.Print("U+2517 (┗) BOX DRAWINGS HEAVY UP AND RIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U2517)

	fmt.Print("U+2518 (┘) BOX DRAWINGS LIGHT UP AND LEFT\n\n")
	fmt.Printf("%z\n", font8x8.U2518)

	fmt.Print("U+2519 (┙) BOX DRAWINGS UP LIGHT AND LEFT HEAVY\n\n")
	fmt.Printf("%z\n", font8x8.U2519)

	fmt.Print("U+251A (┚) BOX DRAWINGS UP HEAVY AND LEFT LIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U251A)

	fmt.Print("U+251B (┛) BOX DRAWINGS HEAVY UP AND LEFT\n\n")
	fmt.Printf("%z\n", font8x8.U251B)

	fmt.Print("U+251C (├) BOX DRAWINGS LIGHT VERTICAL AND RIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U251C)

	fmt.Print("U+251D (┝) BOX DRAWINGS VERTICAL LIGHT AND RIGHT HEAVY\n\n")
	fmt.Printf("%z\n", font8x8.U251D)

	fmt.Print("U+251E (┞) BOX DRAWINGS UP HEAVY AND RIGHT DOWN LIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U251E)

	fmt.Print("U+251F (┟) BOX DRAWINGS DOWN HEAVY AND RIGHT UP LIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U251F)

	fmt.Print("U+2520 (┠) BOX DRAWINGS VERTICAL HEAVY AND RIGHT LIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U2520)

	fmt.Print("U+2521 (┡) BOX DRAWINGS DOWN LIGHT AND RIGHT UP HEAVY\n\n")
	fmt.Printf("%z\n", font8x8.U2521)

	fmt.Print("U+2522 (┢) BOX DRAWINGS UP LIGHT AND RIGHT DOWN HEAVY\n\n")
	fmt.Printf("%z\n", font8x8.U2522)

	fmt.Print("U+2523 (┣) BOX DRAWINGS HEAVY VERTICAL AND RIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U2523)

	fmt.Print("U+2524 (┤) BOX DRAWINGS LIGHT VERTICAL AND LEFT\n\n")
	fmt.Printf("%z\n", font8x8.U2524)

	fmt.Print("U+2525 (┥) BOX DRAWINGS VERTICAL LIGHT AND LEFT HEAVY\n\n")
	fmt.Printf("%z\n", font8x8.U2525)

	fmt.Print("U+2526 (┦) BOX DRAWINGS UP HEAVY AND LEFT DOWN LIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U2526)

	fmt.Print("U+2527 (┧) BOX DRAWINGS DOWN HEAVY AND LEFT UP LIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U2527)

	fmt.Print("U+2528 (┨) BOX DRAWINGS VERTICAL HEAVY AND LEFT LIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U2528)

	fmt.Print("U+2529 (┩) BOX DRAWINGS DOWN LIGHT AND LEFT UP HEAVY\n\n")
	fmt.Printf("%z\n", font8x8.U2529)

	fmt.Print("U+252A (┪) BOX DRAWINGS UP LIGHT AND LEFT DOWN HEAVY\n\n")
	fmt.Printf("%z\n", font8x8.U252A)

	fmt.Print("U+252B (┫) BOX DRAWINGS HEAVY VERTICAL AND LEFT\n\n")
	fmt.Printf("%z\n", font8x8.U252B)

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
	//
	// U+2510 (┐) BOX DRAWINGS LIGHT DOWN AND LEFT
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ████░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	//
	// U+2511 (┑) BOX DRAWINGS DOWN LIGHT AND LEFT HEAVY
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ████░░░░
	// ████░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	//
	// U+2512 (┒) BOX DRAWINGS DOWN HEAVY AND LEFT LIGHT
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// █████░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	//
	// U+2513 (┓) BOX DRAWINGS HEAVY DOWN AND LEFT
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// █████░░░
	// █████░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	//
	// U+2514 (└) BOX DRAWINGS LIGHT UP AND RIGHT
	//
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█████
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+2515 (┕) BOX DRAWINGS UP LIGHT AND RIGHT HEAVY
	//
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█████
	// ░░░█████
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+2516 (┖) BOX DRAWINGS UP HEAVY AND RIGHT LIGHT
	//
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░█████
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+2517 (┗) BOX DRAWINGS HEAVY UP AND RIGHT
	//
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░█████
	// ░░░█████
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+2518 (┘) BOX DRAWINGS LIGHT UP AND LEFT
	//
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ████░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+2519 (┙) BOX DRAWINGS UP LIGHT AND LEFT HEAVY
	//
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ████░░░░
	// ████░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+251A (┚) BOX DRAWINGS UP HEAVY AND LEFT LIGHT
	//
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// █████░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+251B (┛) BOX DRAWINGS HEAVY UP AND LEFT
	//
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// █████░░░
	// █████░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+251C (├) BOX DRAWINGS LIGHT VERTICAL AND RIGHT
	//
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█████
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	//
	// U+251D (┝) BOX DRAWINGS VERTICAL LIGHT AND RIGHT HEAVY
	//
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█████
	// ░░░█████
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	//
	// U+251E (┞) BOX DRAWINGS UP HEAVY AND RIGHT DOWN LIGHT
	//
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░█████
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	//
	// U+251F (┟) BOX DRAWINGS DOWN HEAVY AND RIGHT UP LIGHT
	//
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█████
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	//
	// U+2520 (┠) BOX DRAWINGS VERTICAL HEAVY AND RIGHT LIGHT
	//
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░█████
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	//
	// U+2521 (┡) BOX DRAWINGS DOWN LIGHT AND RIGHT UP HEAVY
	//
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░█████
	// ░░░█████
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	//
	// U+2522 (┢) BOX DRAWINGS UP LIGHT AND RIGHT DOWN HEAVY
	//
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█████
	// ░░░█████
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	//
	// U+2523 (┣) BOX DRAWINGS HEAVY VERTICAL AND RIGHT
	//
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░█████
	// ░░░█████
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	//
	// U+2524 (┤) BOX DRAWINGS LIGHT VERTICAL AND LEFT
	//
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ████░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	//
	// U+2525 (┥) BOX DRAWINGS VERTICAL LIGHT AND LEFT HEAVY
	//
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ████░░░░
	// ████░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	//
	// U+2526 (┦) BOX DRAWINGS UP HEAVY AND LEFT DOWN LIGHT
	//
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// █████░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	//
	// U+2527 (┧) BOX DRAWINGS DOWN HEAVY AND LEFT UP LIGHT
	//
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// █████░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	//
	// U+2528 (┨) BOX DRAWINGS VERTICAL HEAVY AND LEFT LIGHT
	//
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// █████░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	//
	// U+2529 (┩) BOX DRAWINGS DOWN LIGHT AND LEFT UP HEAVY
	//
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// █████░░░
	// █████░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	//
	// U+252A (┪) BOX DRAWINGS UP LIGHT AND LEFT DOWN HEAVY
	//
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// █████░░░
	// █████░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	//
	// U+252B (┫) BOX DRAWINGS HEAVY VERTICAL AND LEFT
	//
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// █████░░░
	// █████░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
}
