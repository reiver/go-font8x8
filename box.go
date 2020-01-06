package font8x8

var (
	U2500 = [8]byte{0x00, 0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00} // U+2500 (─) BOX DRAWINGS LIGHT HORIZONTAL
	U2501 = [8]byte{0x00, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0x00} // U+2501 (━) BOX DRAWINGS HEAVY HORIZONTAL
	U2502 = [8]byte{0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08} // U+2502 (│) BOX DRAWINGS LIGHT VERTICAL
	U2503 = [8]byte{0x18, 0x18, 0x18, 0x18, 0x18, 0x18, 0x18, 0x18} // U+2503 (┃) BOX DRAWINGS HEAVY VERTICAL
	U2504 = [8]byte{0x00, 0x00, 0x00, 0x00, 0xBB, 0x00, 0x00, 0x00} // U+2504 (┄) BOX DRAWINGS LIGHT TRIPLE DASH HORIZONTAL
	U2505 = [8]byte{0x00, 0x00, 0x00, 0xBB, 0xBB, 0x00, 0x00, 0x00} // U+2505 (┅) BOX DRAWINGS HEAVY TRIPLE DASH HORIZONTAL
	U2506 = [8]byte{0x08, 0x00, 0x08, 0x08, 0x08, 0x00, 0x08, 0x08} // U+2506 (┆) BOX DRAWINGS LIGHT TRIPLE DASH VERTICAL
	U2507 = [8]byte{0x18, 0x00, 0x18, 0x18, 0x18, 0x00, 0x18, 0x18} // U+2507 (┇) BOX DRAWINGS HEAVY TRIPLE DASH VERTICAL
	U2508 = [8]byte{0x00, 0x00, 0x00, 0x00, 0x55, 0x00, 0x00, 0x00} // U+2508 (┈) BOX DRAWINGS LIGHT QUADRUPLE DASH HORIZONTAL
	U2509 = [8]byte{0x00, 0x00, 0x00, 0x55, 0x55, 0x00, 0x00, 0x00} // U+2509 (┉) BOX DRAWINGS HEAVY QUADRUPLE DASH HORIZONTAL
	U250A = [8]byte{0x00, 0x08, 0x00, 0x08, 0x00, 0x08, 0x00, 0x08} // U+250A (┊) BOX DRAWINGS LIGHT QUADRUPLE DASH VERTICAL
	U250B = [8]byte{0x00, 0x18, 0x00, 0x18, 0x00, 0x18, 0x00, 0x18} // U+250B (┋) BOX DRAWINGS HEAVY QUADRUPLE DASH VERTICAL
	U250C = [8]byte{0x00, 0x00, 0x00, 0x00, 0xf8, 0x08, 0x08, 0x08} // U+250C (┌) BOX DRAWINGS LIGHT DOWN AND RIGHT
	U250D = [8]byte{0x00, 0x00, 0x00, 0xf8, 0xf8, 0x08, 0x08, 0x08} // U+250D (┍) BOX DRAWINGS DOWN LIGHT AND RIGHT HEAVY
	U250E = [8]byte{0x00, 0x00, 0x00, 0x00, 0xf8, 0x18, 0x18, 0x18} // U+250E (┎) BOX DRAWINGS DOWN HEAVY AND RIGHT LIGHT
	U250F = [8]byte{0x00, 0x00, 0x00, 0xf8, 0xf8, 0x18, 0x18, 0x18} // U+250F (┏) BOX DRAWINGS HEAVY DOWN AND RIGHT
	U2510 = [8]byte{0x00, 0x00, 0x00, 0x00, 0x0f, 0x08, 0x08, 0x08} // U+2510 (┐) BOX DRAWINGS LIGHT DOWN AND LEFT
	U2511 = [8]byte{0x00, 0x00, 0x00, 0x0f, 0x0f, 0x08, 0x08, 0x08} // U+2511 (┑) BOX DRAWINGS DOWN LIGHT AND LEFT HEAVY
	U2512 = [8]byte{0x00, 0x00, 0x00, 0x00, 0x1f, 0x18, 0x18, 0x18} // U+2512 (┒) BOX DRAWINGS DOWN HEAVY AND LEFT LIGHT
	U2513 = [8]byte{0x00, 0x00, 0x00, 0x1f, 0x1f, 0x18, 0x18, 0x18} // U+2513 (┓) BOX DRAWINGS HEAVY DOWN AND LEFT
	U2514 = [8]byte{0x08, 0x08, 0x08, 0x08, 0xf8, 0x00, 0x00, 0x00} // U+2514 (└) BOX DRAWINGS LIGHT UP AND RIGHT
	U2515 = [8]byte{0x08, 0x08, 0x08, 0xf8, 0xf8, 0x00, 0x00, 0x00} // U+2515 (┕) BOX DRAWINGS UP LIGHT AND RIGHT HEAVY
	U2516 = [8]byte{0x18, 0x18, 0x18, 0x18, 0xf8, 0x00, 0x00, 0x00} // U+2516 (┖) BOX DRAWINGS UP HEAVY AND RIGHT LIGHT
	U2517 = [8]byte{0x18, 0x18, 0x18, 0xf8, 0xf8, 0x00, 0x00, 0x00} // U+2517 (┗) BOX DRAWINGS HEAVY UP AND RIGHT
	U2518 = [8]byte{0x08, 0x08, 0x08, 0x08, 0x0f, 0x00, 0x00, 0x00} // U+2518 (┘) BOX DRAWINGS LIGHT UP AND LEFT
	U2519 = [8]byte{0x08, 0x08, 0x08, 0x0f, 0x0f, 0x00, 0x00, 0x00} // U+2519 (┙) BOX DRAWINGS UP LIGHT AND LEFT HEAVY
	U251A = [8]byte{0x18, 0x18, 0x18, 0x18, 0x1f, 0x00, 0x00, 0x00} // U+251A (┚) BOX DRAWINGS UP HEAVY AND LEFT LIGHT
	U251B = [8]byte{0x18, 0x18, 0x18, 0x1f, 0x1f, 0x00, 0x00, 0x00} // U+251B (┛) BOX DRAWINGS HEAVY UP AND LEFT
	U251C = [8]byte{0x08, 0x08, 0x08, 0x08, 0xf8, 0x08, 0x08, 0x08} // U+251C (├) BOX DRAWINGS LIGHT VERTICAL AND RIGHT
	U251D = [8]byte{0x08, 0x08, 0x08, 0xf8, 0xf8, 0x08, 0x08, 0x08} // U+251D (┝) BOX DRAWINGS VERTICAL LIGHT AND RIGHT HEAVY
	U251E = [8]byte{0x18, 0x18, 0x18, 0x18, 0xf8, 0x08, 0x08, 0x08} // U+251E (┞) BOX DRAWINGS UP HEAVY AND RIGHT DOWN LIGHT
	U251F = [8]byte{0x08, 0x08, 0x08, 0x08, 0xf8, 0x18, 0x18, 0x18} // U+251F (┟) BOX DRAWINGS DOWN HEAVY AND RIGHT UP LIGHT
	U2520 = [8]byte{0x18, 0x18, 0x18, 0x18, 0xf8, 0x18, 0x18, 0x18} // U+2520 (┠) BOX DRAWINGS VERTICAL HEAVY AND RIGHT LIGHT
	U2521 = [8]byte{0x18, 0x18, 0x18, 0xf8, 0xf8, 0x08, 0x08, 0x08} // U+2521 (┡) BOX DRAWINGS DOWN LIGHT AND RIGHT UP HEAVY
	U2522 = [8]byte{0x08, 0x08, 0x08, 0xf8, 0xf8, 0x18, 0x18, 0x18} // U+2522 (┢) BOX DRAWINGS UP LIGHT AND RIGHT DOWN HEAVY
	U2523 = [8]byte{0x18, 0x18, 0x18, 0xf8, 0xf8, 0x18, 0x18, 0x18} // U+2523 (┣) BOX DRAWINGS HEAVY VERTICAL AND RIGHT
	U2524 = [8]byte{0x08, 0x08, 0x08, 0x08, 0x0f, 0x08, 0x08, 0x08} // U+2524 (┤) BOX DRAWINGS LIGHT VERTICAL AND LEFT
	U2525 = [8]byte{0x08, 0x08, 0x08, 0x0f, 0x0f, 0x08, 0x08, 0x08} // U+2525 (┥) BOX DRAWINGS VERTICAL LIGHT AND LEFT HEAVY
	U2526 = [8]byte{0x18, 0x18, 0x18, 0x18, 0x1f, 0x08, 0x08, 0x08} // U+2526 (┦) BOX DRAWINGS UP HEAVY AND LEFT DOWN LIGHT
	U2527 = [8]byte{0x08, 0x08, 0x08, 0x08, 0x1f, 0x18, 0x18, 0x18} // U+2527 (┧) BOX DRAWINGS DOWN HEAVY AND LEFT UP LIGHT
	U2528 = [8]byte{0x18, 0x18, 0x18, 0x18, 0x1f, 0x18, 0x18, 0x18} // U+2528 (┨) BOX DRAWINGS VERTICAL HEAVY AND LEFT LIGHT
	U2529 = [8]byte{0x18, 0x18, 0x18, 0x1f, 0x1f, 0x08, 0x08, 0x08} // U+2529 (┩) BOX DRAWINGS DOWN LIGHT AND LEFT UP HEAVY
	U252A = [8]byte{0x08, 0x08, 0x08, 0x1f, 0x1f, 0x18, 0x18, 0x18} // U+252A (┪) BOX DRAWINGS UP LIGHT AND LEFT DOWN HEAVY
	U252B = [8]byte{0x18, 0x18, 0x18, 0x1f, 0x1f, 0x18, 0x18, 0x18} // U+252B (┫) BOX DRAWINGS HEAVY VERTICAL AND LEFT
	U252C = [8]byte{0x00, 0x00, 0x00, 0x00, 0xff, 0x08, 0x08, 0x08} // U+252C (┬) BOX DRAWINGS LIGHT DOWN AND HORIZONTAL
	U252D = [8]byte{0x00, 0x00, 0x00, 0x0f, 0xff, 0x08, 0x08, 0x08} // U+252D (┭) BOX DRAWINGS LEFT HEAVY AND RIGHT DOWN LIGHT
	U252E = [8]byte{0x00, 0x00, 0x00, 0xf8, 0xff, 0x08, 0x08, 0x08} // U+252E (┮) BOX DRAWINGS RIGHT HEAVY AND LEFT DOWN LIGHT
	U252F = [8]byte{0x00, 0x00, 0x00, 0xff, 0xff, 0x08, 0x08, 0x08} // U+252F (┯) BOX DRAWINGS DOWN LIGHT AND HORIZONTAL HEAVY
	U2530 = [8]byte{0x00, 0x00, 0x00, 0x00, 0xff, 0x18, 0x18, 0x18} // U+2530 (┰) BOX DRAWINGS DOWN HEAVY AND HORIZONTAL LIGHT
	U2531 = [8]byte{0x00, 0x00, 0x00, 0x1f, 0xff, 0x18, 0x18, 0x18} // U+2531 (┱) BOX DRAWINGS RIGHT LIGHT AND LEFT DOWN HEAVY
	U2532 = [8]byte{0x00, 0x00, 0x00, 0xf8, 0xff, 0x18, 0x18, 0x18} // U+2532 (┲) BOX DRAWINGS LEFT LIGHT AND RIGHT DOWN HEAVY
	U2533 = [8]byte{0x00, 0x00, 0x00, 0xff, 0xff, 0x18, 0x18, 0x18} // U+2533 (┳) BOX DRAWINGS HEAVY DOWN AND HORIZONTAL
	U2534 = [8]byte{0x08, 0x08, 0x08, 0x08, 0xff, 0x00, 0x00, 0x00} // U+2534 (┴) BOX DRAWINGS LIGHT UP AND HORIZONTAL
	U2535 = [8]byte{0x08, 0x08, 0x08, 0x0f, 0xff, 0x00, 0x00, 0x00} // U+2535 (┵) BOX DRAWINGS LEFT HEAVY AND RIGHT UP LIGHT
	U2536 = [8]byte{0x08, 0x08, 0x08, 0xf8, 0xff, 0x00, 0x00, 0x00} // U+2536 (┶) BOX DRAWINGS RIGHT HEAVY AND LEFT UP LIGHT
	U2537 = [8]byte{0x08, 0x08, 0x08, 0xff, 0xff, 0x00, 0x00, 0x00} // U+2537 (┷) BOX DRAWINGS UP LIGHT AND HORIZONTAL HEAVY
	U2538 = [8]byte{0x18, 0x18, 0x18, 0x18, 0xff, 0x00, 0x00, 0x00} // U+2538 (┸) BOX DRAWINGS UP HEAVY AND HORIZONTAL LIGHT
	U2539 = [8]byte{0x18, 0x18, 0x18, 0x1f, 0xff, 0x00, 0x00, 0x00} // U+2539 (┹) BOX DRAWINGS RIGHT LIGHT AND LEFT UP HEAVY
	U253A = [8]byte{0x18, 0x18, 0x18, 0xf8, 0xff, 0x00, 0x00, 0x00} // U+253A (┺) BOX DRAWINGS LEFT LIGHT AND RIGHT UP HEAVY
	U253B = [8]byte{0x18, 0x18, 0x18, 0xff, 0xff, 0x00, 0x00, 0x00} // U+253B (┻) BOX DRAWINGS HEAVY UP AND HORIZONTAL
	U253C = [8]byte{0x08, 0x08, 0x08, 0x08, 0xff, 0x08, 0x08, 0x08} // U+253C (┼) BOX DRAWINGS LIGHT VERTICAL AND HORIZONTAL
	U253D = [8]byte{0x08, 0x08, 0x08, 0x0f, 0xff, 0x08, 0x08, 0x08} // U+253D (┽) BOX DRAWINGS LEFT HEAVY AND RIGHT VERTICAL LIGHT
	U253E = [8]byte{0x08, 0x08, 0x08, 0xf8, 0xff, 0x08, 0x08, 0x08} // U+253E (┾) BOX DRAWINGS RIGHT HEAVY AND LEFT VERTICAL LIGHT
	U253F = [8]byte{0x08, 0x08, 0x08, 0xff, 0xff, 0x08, 0x08, 0x08} // U+253F (┿) BOX DRAWINGS VERTICAL LIGHT AND HORIZONTAL HEAVY
	U2540 = [8]byte{0x18, 0x18, 0x18, 0x18, 0xff, 0x08, 0x08, 0x08} // U+2540 (╀) BOX DRAWINGS UP HEAVY AND DOWN HORIZONTAL LIGHT
	U2541 = [8]byte{0x08, 0x08, 0x08, 0x08, 0xff, 0x18, 0x18, 0x18} // U+2541 (╁) BOX DRAWINGS DOWN HEAVY AND UP HORIZONTAL LIGHT
	U2542 = [8]byte{0x18, 0x18, 0x18, 0x18, 0xff, 0x18, 0x18, 0x18} // U+2542 (╂) BOX DRAWINGS VERTICAL HEAVY AND HORIZONTAL LIGHT
	U2543 = [8]byte{0x18, 0x18, 0x18, 0x1f, 0xff, 0x08, 0x08, 0x08} // U+2543 (╃) BOX DRAWINGS LEFT UP HEAVY AND RIGHT DOWN LIGHT
	U2544 = [8]byte{0x18, 0x18, 0x18, 0xf8, 0xff, 0x08, 0x08, 0x08} // U+2544 (╄) BOX DRAWINGS RIGHT UP HEAVY AND LEFT DOWN LIGHT
	U2545 = [8]byte{0x08, 0x08, 0x08, 0x1f, 0xff, 0x18, 0x18, 0x18} // U+2545 (╅) BOX DRAWINGS LEFT DOWN HEAVY AND RIGHT UP LIGHT
	U2546 = [8]byte{0x08, 0x08, 0x08, 0xf8, 0xff, 0x18, 0x18, 0x18} // U+2546 (╆) BOX DRAWINGS RIGHT DOWN HEAVY AND LEFT UP LIGHT
	U2547 = [8]byte{0x08, 0x08, 0x08, 0xff, 0xff, 0x18, 0x18, 0x18} // U+2547 (╇) BOX DRAWINGS DOWN LIGHT AND UP HORIZONTAL HEAVY

	U2560 = [8]byte{0x14, 0x14, 0x14, 0xF4, 0x04, 0xF4, 0x14, 0x14} // U+2560 (╠) BOX DRAWINGS DOUBLE VERTICAL AND RIGHT
	U2561 = [8]byte{0x08, 0x08, 0x08, 0x0F, 0x08, 0x0F, 0x08, 0x08} // U+2561 (╡) BOX DRAWINGS VERTICAL SINGLE AND LEFT DOUBLE
	U2562 = [8]byte{0x14, 0x14, 0x14, 0x14, 0x17, 0x14, 0x14, 0x14} // U+2562 (╢) BOX DRAWINGS VERTICAL DOUBLE AND LEFT SINGLE
	U2563 = [8]byte{0x14, 0x14, 0x14, 0x17, 0x10, 0x17, 0x14, 0x14} // U+2563 (╣) BOX DRAWINGS DOUBLE VERTICAL AND LEFT
	U2564 = [8]byte{0x00, 0x00, 0x00, 0xFF, 0x00, 0xFF, 0x08, 0x08} // U+2564 (╤) BOX DRAWINGS DOWN SINGLE AND HORIZONTAL DOUBLE
	U2565 = [8]byte{0x00, 0x00, 0x00, 0x00, 0xFF, 0x14, 0x14, 0x14} // U+2565 (╥) BOX DRAWINGS DOWN DOUBLE AND HORIZONTAL SINGLE
	U2566 = [8]byte{0x00, 0x00, 0x00, 0xFF, 0x00, 0xF7, 0x14, 0x14} // U+2566 (╦) BOX DRAWINGS DOUBLE DOWN AND HORIZONTAL
	U2567 = [8]byte{0x08, 0x08, 0x08, 0xFF, 0x00, 0xFF, 0x00, 0x00} // U+2567 (╧) BOX DRAWINGS UP SINGLE AND HORIZONTAL DOUBLE
	U2568 = [8]byte{0x14, 0x14, 0x14, 0x14, 0xFF, 0x00, 0x00, 0x00} // U+2568 (╨) BOX DRAWINGS UP DOUBLE AND HORIZONTAL SINGLE
	U2569 = [8]byte{0x14, 0x14, 0x14, 0xF7, 0x00, 0xFF, 0x00, 0x00} // U+2569 (╩) BOX DRAWINGS DOUBLE UP AND HORIZONTAL
	U256A = [8]byte{0x08, 0x08, 0x08, 0xFF, 0x08, 0xFF, 0x08, 0x08} // U+256A (╪) BOX DRAWINGS VERTICAL SINGLE AND HORIZONTAL DOUBLE
	U256B = [8]byte{0x14, 0x14, 0x14, 0x14, 0xFF, 0x14, 0x14, 0x14} // U+256B (╫) BOX DRAWINGS VERTICAL DOUBLE AND HORIZONTAL SINGLE
	U256C = [8]byte{0x14, 0x14, 0x14, 0xF7, 0x00, 0xF7, 0x14, 0x14} // U+256C (╬) BOX DRAWINGS DOUBLE VERTICAL AND HORIZONTAL
	U256D = [8]byte{0x00, 0x00, 0x00, 0x00, 0xE0, 0x10, 0x08, 0x08} // U+256D (╭) BOX DRAWINGS LIGHT ARC DOWN AND RIGHT
	U256E = [8]byte{0x00, 0x00, 0x00, 0x00, 0x03, 0x04, 0x08, 0x08} // U+256E (╮) BOX DRAWINGS LIGHT ARC DOWN AND LEFT
	U256F = [8]byte{0x08, 0x08, 0x08, 0x04, 0x03, 0x00, 0x00, 0x00} // U+256F (╯) BOX DRAWINGS LIGHT ARC UP AND LEFT
	U2570 = [8]byte{0x08, 0x08, 0x08, 0x10, 0xE0, 0x00, 0x00, 0x00} // U+2570 (╰) BOX DRAWINGS LIGHT ARC UP AND RIGHT
	U2571 = [8]byte{0x80, 0x40, 0x20, 0x10, 0x08, 0x04, 0x02, 0x01} // U+2571 (╱) BOX DRAWINGS LIGHT DIAGONAL UPPER RIGHT TO LOWER LEFT
	U2572 = [8]byte{0x01, 0x02, 0x04, 0x08, 0x10, 0x20, 0x40, 0x80} // U+2572 (╲) BOX DRAWINGS LIGHT DIAGONAL UPPER LEFT TO LOWER RIGHT
	U2573 = [8]byte{0x81, 0x42, 0x24, 0x18, 0x18, 0x24, 0x42, 0x81} // U+2573 (╳) BOX DRAWINGS LIGHT DIAGONAL CROSS
	U2574 = [8]byte{0x00, 0x00, 0x00, 0x00, 0x0F, 0x00, 0x00, 0x00} // U+2574 (╴) BOX DRAWINGS LIGHT LEFT
	U2575 = [8]byte{0x08, 0x08, 0x08, 0x08, 0x00, 0x00, 0x00, 0x00} // U+2575 (╵) BOX DRAWINGS LIGHT UP
	U2576 = [8]byte{0x00, 0x00, 0x00, 0x00, 0xF8, 0x00, 0x00, 0x00} // U+2576 (╶) BOX DRAWINGS LIGHT RIGHT
	U2577 = [8]byte{0x00, 0x00, 0x00, 0x00, 0x08, 0x08, 0x08, 0x08} // U+2577 (╷) BOX DRAWINGS LIGHT DOWN
	U2578 = [8]byte{0x00, 0x00, 0x00, 0x0F, 0x0F, 0x00, 0x00, 0x00} // U+2578 (╸) BOX DRAWINGS HEAVY LEFT
	U2579 = [8]byte{0x18, 0x18, 0x18, 0x18, 0x00, 0x00, 0x00, 0x00} // U+2579 (╹) BOX DRAWINGS HEAVY UP
	U257A = [8]byte{0x00, 0x00, 0x00, 0xF8, 0xF8, 0x00, 0x00, 0x00} // U+257A (╺) BOX DRAWINGS HEAVY RIGHT
	U257B = [8]byte{0x00, 0x00, 0x00, 0x00, 0x18, 0x18, 0x18, 0x18} // U+257B (╻) BOX DRAWINGS HEAVY DOWN
	U257C = [8]byte{0x00, 0x00, 0x00, 0xF8, 0xFF, 0x00, 0x00, 0x00} // U+257C (╼) BOX DRAWINGS LIGHT LEFT AND HEAVY RIGHT
	U257D = [8]byte{0x08, 0x08, 0x08, 0x08, 0x18, 0x18, 0x18, 0x18} // U+257D (╽) BOX DRAWINGS LIGHT UP AND HEAVY DOWN
	U257E = [8]byte{0x00, 0x00, 0x00, 0x0F, 0xFF, 0x00, 0x00, 0x00} // U+257E (╾) BOX DRAWINGS HEAVY LEFT AND LIGHT RIGHT
	U257F = [8]byte{0x18, 0x18, 0x18, 0x18, 0x08, 0x08, 0x08, 0x08} // U+257F (╿) BOX DRAWINGS HEAVY UP AND LIGHT DOWN
)