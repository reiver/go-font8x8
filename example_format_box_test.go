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

	fmt.Print("U+252C (┬) BOX DRAWINGS LIGHT DOWN AND HORIZONTAL\n\n")
	fmt.Printf("%z\n", font8x8.U252C)

	fmt.Print("U+252D (┭) BOX DRAWINGS LEFT HEAVY AND RIGHT DOWN LIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U252D)

	fmt.Print("U+252E (┮) BOX DRAWINGS RIGHT HEAVY AND LEFT DOWN LIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U252E)

	fmt.Print("U+252F (┯) BOX DRAWINGS DOWN LIGHT AND HORIZONTAL HEAVY\n\n")
	fmt.Printf("%z\n", font8x8.U252F)

	fmt.Print("U+2530 (┰) BOX DRAWINGS DOWN HEAVY AND HORIZONTAL LIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U2530)

	fmt.Print("U+2531 (┱) BOX DRAWINGS RIGHT LIGHT AND LEFT DOWN HEAVY\n\n")
	fmt.Printf("%z\n", font8x8.U2531)

	fmt.Print("U+2532 (┲) BOX DRAWINGS LEFT LIGHT AND RIGHT DOWN HEAVY\n\n")
	fmt.Printf("%z\n", font8x8.U2532)

	fmt.Print("U+2533 (┳) BOX DRAWINGS HEAVY DOWN AND HORIZONTAL\n\n")
	fmt.Printf("%z\n", font8x8.U2533)

	fmt.Print("U+2534 (┴) BOX DRAWINGS LIGHT UP AND HORIZONTAL\n\n")
	fmt.Printf("%z\n", font8x8.U2534)

	fmt.Print("U+2535 (┵) BOX DRAWINGS LEFT HEAVY AND RIGHT UP LIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U2535)

	fmt.Print("U+2536 (┶) BOX DRAWINGS RIGHT HEAVY AND LEFT UP LIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U2536)

	fmt.Print("U+2537 (┷) BOX DRAWINGS UP LIGHT AND HORIZONTAL HEAVY\n\n")
	fmt.Printf("%z\n", font8x8.U2537)

	fmt.Print("U+2538 (┸) BOX DRAWINGS UP HEAVY AND HORIZONTAL LIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U2538)

	fmt.Print("U+2539 (┹) BOX DRAWINGS RIGHT LIGHT AND LEFT UP HEAVY\n\n")
	fmt.Printf("%z\n", font8x8.U2539)

	fmt.Print("U+253A (┺) BOX DRAWINGS LEFT LIGHT AND RIGHT UP HEAVY\n\n")
	fmt.Printf("%z\n", font8x8.U253A)

	fmt.Print("U+253B (┻) BOX DRAWINGS HEAVY UP AND HORIZONTAL\n\n")
	fmt.Printf("%z\n", font8x8.U253B)

	fmt.Print("U+253C (┼) BOX DRAWINGS LIGHT VERTICAL AND HORIZONTAL\n\n")
	fmt.Printf("%z\n", font8x8.U253C)

	fmt.Print("U+253D (┽) BOX DRAWINGS LEFT HEAVY AND RIGHT VERTICAL LIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U253D)

	fmt.Print("U+253E (┾) BOX DRAWINGS RIGHT HEAVY AND LEFT VERTICAL LIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U253E)

	fmt.Print("U+253F (┿) BOX DRAWINGS VERTICAL LIGHT AND HORIZONTAL HEAVY\n\n")
	fmt.Printf("%z\n", font8x8.U253F)

	fmt.Print("U+2540 (╀) BOX DRAWINGS UP HEAVY AND DOWN HORIZONTAL LIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U2540)

	fmt.Print("U+2541 (╁) BOX DRAWINGS DOWN HEAVY AND UP HORIZONTAL LIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U2541)

	fmt.Print("U+2542 (╂) BOX DRAWINGS VERTICAL HEAVY AND HORIZONTAL LIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U2542)

	fmt.Print("U+2543 (╃) BOX DRAWINGS LEFT UP HEAVY AND RIGHT DOWN LIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U2543)

	fmt.Print("U+2544 (╄) BOX DRAWINGS RIGHT UP HEAVY AND LEFT DOWN LIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U2544)

	fmt.Print("U+2545 (╅) BOX DRAWINGS LEFT DOWN HEAVY AND RIGHT UP LIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U2545)

	fmt.Print("U+2546 (╆) BOX DRAWINGS RIGHT DOWN HEAVY AND LEFT UP LIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U2546)

	fmt.Print("U+2547 (╈) BOX DRAWINGS DOWN LIGHT AND UP HORIZONTAL HEAVY\n\n")
	fmt.Printf("%z\n", font8x8.U2547)

	fmt.Print("U+2548 (╇) BOX DRAWINGS UP LIGHT AND DOWN HORIZONTAL HEAVY\n\n")
	fmt.Printf("%z\n", font8x8.U2548)

	fmt.Print("U+2549 (╉) BOX DRAWINGS RIGHT LIGHT AND LEFT VERTICAL HEAVY\n\n")
	fmt.Printf("%z\n", font8x8.U2549)

	fmt.Print("U+254A (╊) BOX DRAWINGS LEFT LIGHT AND RIGHT VERTICAL HEAVY\n\n")
	fmt.Printf("%z\n", font8x8.U254A)

	fmt.Print("U+254B (╋) BOX DRAWINGS HEAVY VERTICAL AND HORIZONTAL\n\n")
	fmt.Printf("%z\n", font8x8.U254B)

	fmt.Print("U+254C (╌) BOX DRAWINGS LIGHT DOUBLE DASH HORIZONTAL\n\n")
	fmt.Printf("%z\n", font8x8.U254C)

	fmt.Print("U+254D (╍) BOX DRAWINGS HEAVY DOUBLE DASH HORIZONTAL\n\n")
	fmt.Printf("%z\n", font8x8.U254D)

	fmt.Print("U+254E (╎) BOX DRAWINGS LIGHT DOUBLE DASH VERTICAL\n\n")
	fmt.Printf("%z\n", font8x8.U254E)

	fmt.Print("U+254F (╏) BOX DRAWINGS HEAVY DOUBLE DASH VERTICAL\n\n")
	fmt.Printf("%z\n", font8x8.U254F)

	fmt.Print("U+2550 (═) BOX DRAWINGS DOUBLE HORIZONTAL\n\n")
	fmt.Printf("%z\n", font8x8.U2550)

	fmt.Print("U+2551 (║) BOX DRAWINGS DOUBLE VERTICAL\n\n")
	fmt.Printf("%z\n", font8x8.U2551)

	fmt.Print("U+2552 (╒) BOX DRAWINGS DOWN SINGLE AND RIGHT DOUBLE\n\n")
	fmt.Printf("%z\n", font8x8.U2552)

	fmt.Print("U+2553 (╓) BOX DRAWINGS DOWN DOUBLE AND RIGHT SINGLE\n\n")
	fmt.Printf("%z\n", font8x8.U2553)

	fmt.Print("U+2554 (╔) BOX DRAWINGS DOUBLE DOWN AND RIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U2554)

	fmt.Print("U+2555 (╕) BOX DRAWINGS DOWN SINGLE AND LEFT DOUBLE\n\n")
	fmt.Printf("%z\n", font8x8.U2555)

	fmt.Print("U+2556 (╖) BOX DRAWINGS DOWN DOUBLE AND LEFT SINGLE\n\n")
	fmt.Printf("%z\n", font8x8.U2556)

	fmt.Print("U+2557 (╗) BOX DRAWINGS DOUBLE DOWN AND LEFT\n\n")
	fmt.Printf("%z\n", font8x8.U2557)

	fmt.Print("U+2558 (╘) BOX DRAWINGS UP SINGLE AND RIGHT DOUBLE\n\n")
	fmt.Printf("%z\n", font8x8.U2558)

	fmt.Print("U+2559 (╙) BOX DRAWINGS UP DOUBLE AND RIGHT SINGLE\n\n")
	fmt.Printf("%z\n", font8x8.U2559)

	fmt.Print("U+255A (╚) BOX DRAWINGS DOUBLE UP AND RIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U255A)

	fmt.Print("U+255B (╛) BOX DRAWINGS UP SINGLE AND LEFT DOUBLE\n\n")
	fmt.Printf("%z\n", font8x8.U255B)

	fmt.Print("U+255C (╜) BOX DRAWINGS UP DOUBLE AND LEFT SINGLE\n\n")
	fmt.Printf("%z\n", font8x8.U255C)

	fmt.Print("U+255D (╝) BOX DRAWINGS DOUBLE UP AND LEFT\n\n")
	fmt.Printf("%z\n", font8x8.U255D)

	fmt.Print("U+255E (╞) BOX DRAWINGS VERTICAL SINGLE AND RIGHT DOUBLE\n\n")
	fmt.Printf("%z\n", font8x8.U255E)

	fmt.Print("U+255F (╟) BOX DRAWINGS VERTICAL DOUBLE AND RIGHT SINGLE\n\n")
	fmt.Printf("%z\n", font8x8.U255F)

	fmt.Print("U+2560 (╠) BOX DRAWINGS DOUBLE VERTICAL AND RIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U2560)

	fmt.Print("U+2561 (╡) BOX DRAWINGS VERTICAL SINGLE AND LEFT DOUBLE\n\n")
	fmt.Printf("%z\n", font8x8.U2561)

	fmt.Print("U+2562 (╢) BOX DRAWINGS VERTICAL DOUBLE AND LEFT SINGLE\n\n")
	fmt.Printf("%z\n", font8x8.U2562)

	fmt.Print("U+2563 (╣) BOX DRAWINGS DOUBLE VERTICAL AND LEFT\n\n")
	fmt.Printf("%z\n", font8x8.U2563)

	fmt.Print("U+2564 (╤) BOX DRAWINGS DOWN SINGLE AND HORIZONTAL DOUBLE\n\n")
	fmt.Printf("%z\n", font8x8.U2564)

	fmt.Print("U+2565 (╥) BOX DRAWINGS DOWN DOUBLE AND HORIZONTAL SINGLE\n\n")
	fmt.Printf("%z\n", font8x8.U2565)

	fmt.Print("U+2566 (╦) BOX DRAWINGS DOUBLE DOWN AND HORIZONTAL\n\n")
	fmt.Printf("%z\n", font8x8.U2566)

	fmt.Print("U+2567 (╧) BOX DRAWINGS UP SINGLE AND HORIZONTAL DOUBLE\n\n")
	fmt.Printf("%z\n", font8x8.U2567)

	fmt.Print("U+2568 (╨) BOX DRAWINGS UP DOUBLE AND HORIZONTAL SINGLE\n\n")
	fmt.Printf("%z\n", font8x8.U2568)

	fmt.Print("U+2569 (╩) BOX DRAWINGS DOUBLE UP AND HORIZONTAL\n\n")
	fmt.Printf("%z\n", font8x8.U2569)

	fmt.Print("U+256A (╪) BOX DRAWINGS VERTICAL SINGLE AND HORIZONTAL DOUBLE\n\n")
	fmt.Printf("%z\n", font8x8.U256A)

	fmt.Print("U+256B (╫) BOX DRAWINGS VERTICAL DOUBLE AND HORIZONTAL SINGLE\n\n")
	fmt.Printf("%z\n", font8x8.U256B)

	fmt.Print("U+256C (╬) BOX DRAWINGS DOUBLE VERTICAL AND HORIZONTAL\n\n")
	fmt.Printf("%z\n", font8x8.U256C)

	fmt.Print("U+256D (╭) BOX DRAWINGS LIGHT ARC DOWN AND RIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U256D)

	fmt.Print("U+256E (╮) BOX DRAWINGS LIGHT ARC DOWN AND LEFT\n\n")
	fmt.Printf("%z\n", font8x8.U256E)

	fmt.Print("U+256F (╯) BOX DRAWINGS LIGHT ARC UP AND LEFT\n\n")
	fmt.Printf("%z\n", font8x8.U256F)

	fmt.Print("U+2570 (╰) BOX DRAWINGS LIGHT ARC UP AND RIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U2570)

	fmt.Print("U+2571 (╱) BOX DRAWINGS LIGHT DIAGONAL UPPER RIGHT TO LOWER LEFT\n\n")
	fmt.Printf("%z\n", font8x8.U2571)

	fmt.Print("U+2572 (╲) BOX DRAWINGS LIGHT DIAGONAL UPPER LEFT TO LOWER RIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U2572)

	fmt.Print("U+2573 (╳) BOX DRAWINGS LIGHT DIAGONAL CROSS\n\n")
	fmt.Printf("%z\n", font8x8.U2573)

	fmt.Print("U+2574 (╴) BOX DRAWINGS LIGHT LEFT\n\n")
	fmt.Printf("%z\n", font8x8.U2574)

	fmt.Print("U+2575 (╵) BOX DRAWINGS LIGHT UP\n\n")
	fmt.Printf("%z\n", font8x8.U2575)

	fmt.Print("U+2576 (╶) BOX DRAWINGS LIGHT RIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U2576)

	fmt.Print("U+2577 (╷) BOX DRAWINGS LIGHT DOWN\n\n")
	fmt.Printf("%z\n", font8x8.U2577)

	fmt.Print("U+2578 (╸) BOX DRAWINGS HEAVY LEFT\n\n")
	fmt.Printf("%z\n", font8x8.U2578)

	fmt.Print("U+2579 (╹) BOX DRAWINGS HEAVY UP\n\n")
	fmt.Printf("%z\n", font8x8.U2579)

	fmt.Print("U+257A (╺) BOX DRAWINGS HEAVY RIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U257A)

	fmt.Print("U+257B (╻) BOX DRAWINGS HEAVY DOWN\n\n")
	fmt.Printf("%z\n", font8x8.U257B)

	fmt.Print("U+257C (╼) BOX DRAWINGS LIGHT LEFT AND HEAVY RIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U257C)

	fmt.Print("U+257D (╽) BOX DRAWINGS LIGHT UP AND HEAVY DOWN\n\n")
	fmt.Printf("%z\n", font8x8.U257D)

	fmt.Print("U+257E (╾) BOX DRAWINGS HEAVY LEFT AND LIGHT RIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U257E)

	fmt.Print("U+257F (╿) BOX DRAWINGS HEAVY UP AND LIGHT DOWN\n\n")
	fmt.Printf("%z\n", font8x8.U257F)

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
	//
	// U+252C (┬) BOX DRAWINGS LIGHT DOWN AND HORIZONTAL
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ████████
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	//
	// U+252D (┭) BOX DRAWINGS LEFT HEAVY AND RIGHT DOWN LIGHT
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ████░░░░
	// ████████
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	//
	// U+252E (┮) BOX DRAWINGS RIGHT HEAVY AND LEFT DOWN LIGHT
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░█████
	// ████████
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	//
	// U+252F (┯) BOX DRAWINGS DOWN LIGHT AND HORIZONTAL HEAVY
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ████████
	// ████████
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	//
	// U+2530 (┰) BOX DRAWINGS DOWN HEAVY AND HORIZONTAL LIGHT
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ████████
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	//
	// U+2531 (┱) BOX DRAWINGS RIGHT LIGHT AND LEFT DOWN HEAVY
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// █████░░░
	// ████████
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	//
	// U+2532 (┲) BOX DRAWINGS LEFT LIGHT AND RIGHT DOWN HEAVY
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░█████
	// ████████
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	//
	// U+2533 (┳) BOX DRAWINGS HEAVY DOWN AND HORIZONTAL
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ████████
	// ████████
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	//
	// U+2534 (┴) BOX DRAWINGS LIGHT UP AND HORIZONTAL
	//
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ████████
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+2535 (┵) BOX DRAWINGS LEFT HEAVY AND RIGHT UP LIGHT
	//
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ████░░░░
	// ████████
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+2536 (┶) BOX DRAWINGS RIGHT HEAVY AND LEFT UP LIGHT
	//
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█████
	// ████████
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+2537 (┷) BOX DRAWINGS UP LIGHT AND HORIZONTAL HEAVY
	//
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ████████
	// ████████
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+2538 (┸) BOX DRAWINGS UP HEAVY AND HORIZONTAL LIGHT
	//
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// ████████
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+2539 (┹) BOX DRAWINGS RIGHT LIGHT AND LEFT UP HEAVY
	//
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// █████░░░
	// ████████
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+253A (┺) BOX DRAWINGS LEFT LIGHT AND RIGHT UP HEAVY
	//
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░█████
	// ████████
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+253B (┻) BOX DRAWINGS HEAVY UP AND HORIZONTAL
	//
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// ████████
	// ████████
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+253C (┼) BOX DRAWINGS LIGHT VERTICAL AND HORIZONTAL
	//
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ████████
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	//
	// U+253D (┽) BOX DRAWINGS LEFT HEAVY AND RIGHT VERTICAL LIGHT
	//
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ████░░░░
	// ████████
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	//
	// U+253E (┾) BOX DRAWINGS RIGHT HEAVY AND LEFT VERTICAL LIGHT
	//
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█████
	// ████████
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	//
	// U+253F (┿) BOX DRAWINGS VERTICAL LIGHT AND HORIZONTAL HEAVY
	//
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ████████
	// ████████
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	//
	// U+2540 (╀) BOX DRAWINGS UP HEAVY AND DOWN HORIZONTAL LIGHT
	//
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// ████████
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	//
	// U+2541 (╁) BOX DRAWINGS DOWN HEAVY AND UP HORIZONTAL LIGHT
	//
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ████████
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	//
	// U+2542 (╂) BOX DRAWINGS VERTICAL HEAVY AND HORIZONTAL LIGHT
	//
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// ████████
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	//
	// U+2543 (╃) BOX DRAWINGS LEFT UP HEAVY AND RIGHT DOWN LIGHT
	//
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// █████░░░
	// ████████
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	//
	// U+2544 (╄) BOX DRAWINGS RIGHT UP HEAVY AND LEFT DOWN LIGHT
	//
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░█████
	// ████████
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	//
	// U+2545 (╅) BOX DRAWINGS LEFT DOWN HEAVY AND RIGHT UP LIGHT
	//
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// █████░░░
	// ████████
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	//
	// U+2546 (╆) BOX DRAWINGS RIGHT DOWN HEAVY AND LEFT UP LIGHT
	//
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█████
	// ████████
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	//
	// U+2547 (╈) BOX DRAWINGS DOWN LIGHT AND UP HORIZONTAL HEAVY
	//
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// ████████
	// ████████
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	//
	// U+2548 (╇) BOX DRAWINGS UP LIGHT AND DOWN HORIZONTAL HEAVY
	//
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ████████
	// ████████
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	//
	// U+2549 (╉) BOX DRAWINGS RIGHT LIGHT AND LEFT VERTICAL HEAVY
	//
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// █████░░░
	// ████████
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	//
	// U+254A (╊) BOX DRAWINGS LEFT LIGHT AND RIGHT VERTICAL HEAVY
	//
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░█████
	// ████████
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	//
	// U+254B (╋) BOX DRAWINGS HEAVY VERTICAL AND HORIZONTAL
	//
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// ████████
	// ████████
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	//
	// U+254C (╌) BOX DRAWINGS LIGHT DOUBLE DASH HORIZONTAL
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ███░░███
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+254D (╍) BOX DRAWINGS HEAVY DOUBLE DASH HORIZONTAL
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ███░░███
	// ███░░███
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+254E (╎) BOX DRAWINGS LIGHT DOUBLE DASH VERTICAL
	//
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	//
	// U+254F (╏) BOX DRAWINGS HEAVY DOUBLE DASH VERTICAL
	//
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	//
	// U+2550 (═) BOX DRAWINGS DOUBLE HORIZONTAL
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ████████
	// ░░░░░░░░
	// ████████
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+2551 (║) BOX DRAWINGS DOUBLE VERTICAL
	//
	// ░░█░█░░░
	// ░░█░█░░░
	// ░░█░█░░░
	// ░░█░█░░░
	// ░░█░█░░░
	// ░░█░█░░░
	// ░░█░█░░░
	// ░░█░█░░░
	//
	// U+2552 (╒) BOX DRAWINGS DOWN SINGLE AND RIGHT DOUBLE
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░█████
	// ░░░█░░░░
	// ░░░█████
	// ░░░█░░░░
	// ░░░█░░░░
	//
	// U+2553 (╓) BOX DRAWINGS DOWN DOUBLE AND RIGHT SINGLE
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░██████
	// ░░█░█░░░
	// ░░█░█░░░
	// ░░█░█░░░
	//
	// U+2554 (╔) BOX DRAWINGS DOUBLE DOWN AND RIGHT
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░██████
	// ░░█░░░░░
	// ░░█░████
	// ░░█░█░░░
	// ░░█░█░░░
	//
	// U+2555 (╕) BOX DRAWINGS DOWN SINGLE AND LEFT DOUBLE
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ████░░░░
	// ░░░█░░░░
	// ████░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	//
	// U+2556 (╖) BOX DRAWINGS DOWN DOUBLE AND LEFT SINGLE
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// █████░░░
	// ░░█░█░░░
	// ░░█░█░░░
	// ░░█░█░░░
	//
	// U+2557 (╗) BOX DRAWINGS DOUBLE DOWN AND LEFT
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// █████░░░
	// ░░░░█░░░
	// ███░█░░░
	// ░░█░█░░░
	// ░░█░█░░░
	//
	// U+2558 (╘) BOX DRAWINGS UP SINGLE AND RIGHT DOUBLE
	//
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█████
	// ░░░█░░░░
	// ░░░█████
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+2559 (╙) BOX DRAWINGS UP DOUBLE AND RIGHT SINGLE
	//
	// ░░█░█░░░
	// ░░█░█░░░
	// ░░█░█░░░
	// ░░█░█░░░
	// ░░██████
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+255A (╚) BOX DRAWINGS DOUBLE UP AND RIGHT
	//
	// ░░█░█░░░
	// ░░█░█░░░
	// ░░█░█░░░
	// ░░█░████
	// ░░█░░░░░
	// ░░██████
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+255B (╛) BOX DRAWINGS UP SINGLE AND LEFT DOUBLE
	//
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ████░░░░
	// ░░░█░░░░
	// ████░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+255C (╜) BOX DRAWINGS UP DOUBLE AND LEFT SINGLE
	//
	// ░░█░█░░░
	// ░░█░█░░░
	// ░░█░█░░░
	// ░░█░█░░░
	// █████░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+255D (╝) BOX DRAWINGS DOUBLE UP AND LEFT
	//
	// ░░█░█░░░
	// ░░█░█░░░
	// ░░█░█░░░
	// ███░█░░░
	// ░░░░█░░░
	// █████░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+255E (╞) BOX DRAWINGS VERTICAL SINGLE AND RIGHT DOUBLE
	//
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█████
	// ░░░█░░░░
	// ░░░█████
	// ░░░█░░░░
	// ░░░█░░░░
	//
	// U+255F (╟) BOX DRAWINGS VERTICAL DOUBLE AND RIGHT SINGLE
	//
	// ░░█░█░░░
	// ░░█░█░░░
	// ░░█░█░░░
	// ░░█░█░░░
	// ░░█░████
	// ░░█░█░░░
	// ░░█░█░░░
	// ░░█░█░░░
	//
	// U+2560 (╠) BOX DRAWINGS DOUBLE VERTICAL AND RIGHT
	//
	// ░░█░█░░░
	// ░░█░█░░░
	// ░░█░█░░░
	// ░░█░████
	// ░░█░░░░░
	// ░░█░████
	// ░░█░█░░░
	// ░░█░█░░░
	//
	// U+2561 (╡) BOX DRAWINGS VERTICAL SINGLE AND LEFT DOUBLE
	//
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ████░░░░
	// ░░░█░░░░
	// ████░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	//
	// U+2562 (╢) BOX DRAWINGS VERTICAL DOUBLE AND LEFT SINGLE
	//
	// ░░█░█░░░
	// ░░█░█░░░
	// ░░█░█░░░
	// ░░█░█░░░
	// ███░█░░░
	// ░░█░█░░░
	// ░░█░█░░░
	// ░░█░█░░░
	//
	// U+2563 (╣) BOX DRAWINGS DOUBLE VERTICAL AND LEFT
	//
	// ░░█░█░░░
	// ░░█░█░░░
	// ░░█░█░░░
	// ███░█░░░
	// ░░░░█░░░
	// ███░█░░░
	// ░░█░█░░░
	// ░░█░█░░░
	//
	// U+2564 (╤) BOX DRAWINGS DOWN SINGLE AND HORIZONTAL DOUBLE
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ████████
	// ░░░░░░░░
	// ████████
	// ░░░█░░░░
	// ░░░█░░░░
	//
	// U+2565 (╥) BOX DRAWINGS DOWN DOUBLE AND HORIZONTAL SINGLE
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ████████
	// ░░█░█░░░
	// ░░█░█░░░
	// ░░█░█░░░
	//
	// U+2566 (╦) BOX DRAWINGS DOUBLE DOWN AND HORIZONTAL
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ████████
	// ░░░░░░░░
	// ███░████
	// ░░█░█░░░
	// ░░█░█░░░
	//
	// U+2567 (╧) BOX DRAWINGS UP SINGLE AND HORIZONTAL DOUBLE
	//
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ████████
	// ░░░░░░░░
	// ████████
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+2568 (╨) BOX DRAWINGS UP DOUBLE AND HORIZONTAL SINGLE
	//
	// ░░█░█░░░
	// ░░█░█░░░
	// ░░█░█░░░
	// ░░█░█░░░
	// ████████
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+2569 (╩) BOX DRAWINGS DOUBLE UP AND HORIZONTAL
	//
	// ░░█░█░░░
	// ░░█░█░░░
	// ░░█░█░░░
	// ███░████
	// ░░░░░░░░
	// ████████
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+256A (╪) BOX DRAWINGS VERTICAL SINGLE AND HORIZONTAL DOUBLE
	//
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ████████
	// ░░░█░░░░
	// ████████
	// ░░░█░░░░
	// ░░░█░░░░
	//
	// U+256B (╫) BOX DRAWINGS VERTICAL DOUBLE AND HORIZONTAL SINGLE
	//
	// ░░█░█░░░
	// ░░█░█░░░
	// ░░█░█░░░
	// ░░█░█░░░
	// ████████
	// ░░█░█░░░
	// ░░█░█░░░
	// ░░█░█░░░
	//
	// U+256C (╬) BOX DRAWINGS DOUBLE VERTICAL AND HORIZONTAL
	//
	// ░░█░█░░░
	// ░░█░█░░░
	// ░░█░█░░░
	// ███░████
	// ░░░░░░░░
	// ███░████
	// ░░█░█░░░
	// ░░█░█░░░
	//
	// U+256D (╭) BOX DRAWINGS LIGHT ARC DOWN AND RIGHT
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░███
	// ░░░░█░░░
	// ░░░█░░░░
	// ░░░█░░░░
	//
	// U+256E (╮) BOX DRAWINGS LIGHT ARC DOWN AND LEFT
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ██░░░░░░
	// ░░█░░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	//
	// U+256F (╯) BOX DRAWINGS LIGHT ARC UP AND LEFT
	//
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░█░░░░░
	// ██░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+2570 (╰) BOX DRAWINGS LIGHT ARC UP AND RIGHT
	//
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░░█░░░
	// ░░░░░███
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+2571 (╱) BOX DRAWINGS LIGHT DIAGONAL UPPER RIGHT TO LOWER LEFT
	//
	// ░░░░░░░█
	// ░░░░░░█░
	// ░░░░░█░░
	// ░░░░█░░░
	// ░░░█░░░░
	// ░░█░░░░░
	// ░█░░░░░░
	// █░░░░░░░
	//
	// U+2572 (╲) BOX DRAWINGS LIGHT DIAGONAL UPPER LEFT TO LOWER RIGHT
	//
	// █░░░░░░░
	// ░█░░░░░░
	// ░░█░░░░░
	// ░░░█░░░░
	// ░░░░█░░░
	// ░░░░░█░░
	// ░░░░░░█░
	// ░░░░░░░█
	//
	// U+2573 (╳) BOX DRAWINGS LIGHT DIAGONAL CROSS
	//
	// █░░░░░░█
	// ░█░░░░█░
	// ░░█░░█░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░█░░█░░
	// ░█░░░░█░
	// █░░░░░░█
	//
	// U+2574 (╴) BOX DRAWINGS LIGHT LEFT
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ████░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+2575 (╵) BOX DRAWINGS LIGHT UP
	//
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+2576 (╶) BOX DRAWINGS LIGHT RIGHT
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░█████
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+2577 (╷) BOX DRAWINGS LIGHT DOWN
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	//
	// U+2578 (╸) BOX DRAWINGS HEAVY LEFT
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ████░░░░
	// ████░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+2579 (╹) BOX DRAWINGS HEAVY UP
	//
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+257A (╺) BOX DRAWINGS HEAVY RIGHT
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░█████
	// ░░░█████
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+257B (╻) BOX DRAWINGS HEAVY DOWN
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	//
	// U+257C (╼) BOX DRAWINGS LIGHT LEFT AND HEAVY RIGHT
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░█████
	// ████████
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+257D (╽) BOX DRAWINGS LIGHT UP AND HEAVY DOWN
	//
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	//
	// U+257E (╾) BOX DRAWINGS HEAVY LEFT AND LIGHT RIGHT
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ████░░░░
	// ████████
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+257F (╿) BOX DRAWINGS HEAVY UP AND LIGHT DOWN
	//
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
	// ░░░█░░░░
}
