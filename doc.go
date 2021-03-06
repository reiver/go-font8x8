/*

Package font8x8 provides a set of 8×8 fonts for Unicode characters.

Fprintf, Printf, Sprintf Support

This package has support for Go's built-in "fmt" package.

The verb:

	%z     prints the font in block characters.

For example:

	fmt.Printf("%z", font8x8.U0052)

Would output:

	██████░░
	░██░░██░
	░██░░██░
	░█████░░
	░██░██░░
	░██░░██░
	███░░██░
	░░░░░░░░

Image Draw Support

This package has support for Go's built-in "image/draw" package.

In particular, a font8x8.Type fits the image.Image interface, and thus can be used as a _source_ image when drawing.

See the following article for more information on how to use Go's built-in "image/draw" package:

https://blog.golang.org/go-imagedraw-package

Font Format

This pacage provides 8×8 fonts. Each glyph is stored in a [8]byte array.

An example may make the format more clear.

Here is LATIN CAPITAL LETTER A ('A') — U+0041:

	font8x8.U0041 = [8]byte{0x0C, 0x1E, 0x33, 0x33, 0x3F, 0x33, 0x33, 0x00}

(Go doesn't have binary literals, but let's pretend it does so then) this the same [8]byte using binary literals would be:

	font8x8.U0041 = [8]byte{0b00001100, 0b00011110, 0b00110011, 0b00110011, 0b00111111, 0b00110011, 0b00110011, 0b00000000}

And formatting that a bit differently (for clarity) we get:

	font8x8.U0041 = [8]byte{
		0b00001100,
		0b00011110,
		0b00110011,
		0b00110011,
		0b00111111,
		0b00110011,
		0b00110011,
		0b00000000
	}

This would correspond to:

	░░██░░░░
	░████░░░
	██░░██░░
	██░░██░░
	██████░░
	██░░██░░
	██░░██░░
	░░░░░░░░

I.e., zero (0) is empty space, and one (1) is character.

Credits

These font were (originally) created by Marcel Sondaar and released by him into the public domain.

	; Summary: font8_8.asm
	; 8x8 monochrome bitmap fonts for rendering
	;
	; Author:
	;     Marcel Sondaar
	;     International Business Machines (public domain VGA fonts)
	;
	; License:
	;     Public Domain
	;

Originally fetched from: http://dimensionalrift.homelinux.net/combuster/mos3/?p=viewsource&file=/modules/gfx/font8_8.asm
*/
package font8x8


