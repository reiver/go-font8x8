package font8x8_test

import (
	"github.com/reiver/go-font8x8"

	"fmt"
	"image"
	"image/draw"
)

func ExampleType_draw_U0043() {

	// We create a byte array that will store the destination image.
	//
	// This byte array is named: ‘dstMemory’.
	//
	// It has a dimesions of: width × height = ‘dstWidth’ × ‘dstHeight’ = 32×16.
	//
	// Each pixel is ‘dstDepth’ bytes — 4 bytes — (rather than just 1 byte)
	// because each pixel is an RGBA value.
	//
	// Therefore the total number of bytes it has is:
	//
	// ‘dstWidth’ * ‘dstHeight’ * ‘dstDepth’ = 32*16*4 = 2048

	const dstWidth  = 14
	const dstHeight = 12
	const dstDepth  = 4  // RGBA

	var dstMemory [dstWidth*dstHeight*dstDepth]byte

	// Go has built-in "image" & "image/draw" packages that provide tools for working with images.
	//
	// We are going to use these built-in Go packages.
	//
	// To do that, we need to make our byte array — ‘dstMemory’ — to fit the ‘draw.Image’ interface.
	//
	// To do that we will wrap it in a ‘image.NRGBA’.
	//
	// The wrapped ‘dstMemory’ will be called: ‘dst’.

	var dst image.NRGBA = image.NRGBA{
		Pix: dstMemory[:],
		Stride: dstWidth*dstDepth,
		Rect: image.Rectangle{
			Min: image.Point{
				X:0,
				Y:0,
			},
			Max: image.Point{
				X:(dstWidth-1),
				Y:(dstHeight-1),
			},
		},
	}

	// Here we get our 8x8 font.
	//
	// A ‘font8x8.Type’ already fits the ‘image.Image’ interface, so we do not need to wrap it
	// to make it work with ‘draw.Draw()’.

	var src font8x8.Type = font8x8.Rune('\u0043')

	const fontWidth  = 8
	const fontHeight = 8

	// Now we draw the 8x8 font — which we have in ‘src’ — to ‘dst’
	// (which will cause it to be drawn on ‘dstMemory’).
	//
	// We will draw it at (x,y) = (2,4).

	const x = 2
	const y = 4

	rect := image.Rectangle{
		Min: image.Point{
			X:x,
			Y:x,
		},
		Max: image.Point{
			X:x+(fontWidth),
			Y:y+(fontHeight),
		},
	}


	draw.Draw(&dst, rect, src, image.ZP, draw.Src)

	// Now we will output what is ‘dstMemory’ to show that the drawing worked!

	for offset, value := range dstMemory {
		if 0 == offset % (dstWidth*dstDepth) {
			fmt.Println()
		} else if 0 == offset % 4 {
			fmt.Print(" ")
		}

		fmt.Printf("%02X", value)
	}

	// Output:
	// 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000
	// 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000
	// 00000000 00000000 000000FF 000000FF FFFFFFFF FFFFFFFF FFFFFFFF FFFFFFFF 000000FF 00000000 00000000 00000000 00000000 00000000
	// 00000000 00000000 000000FF FFFFFFFF FFFFFFFF 000000FF 000000FF FFFFFFFF FFFFFFFF 00000000 00000000 00000000 00000000 00000000
	// 00000000 00000000 FFFFFFFF FFFFFFFF 000000FF 000000FF 000000FF 000000FF 000000FF 00000000 00000000 00000000 00000000 00000000
	// 00000000 00000000 FFFFFFFF FFFFFFFF 000000FF 000000FF 000000FF 000000FF 000000FF 00000000 00000000 00000000 00000000 00000000
	// 00000000 00000000 FFFFFFFF FFFFFFFF 000000FF 000000FF 000000FF 000000FF 000000FF 00000000 00000000 00000000 00000000 00000000
	// 00000000 00000000 000000FF FFFFFFFF FFFFFFFF 000000FF 000000FF FFFFFFFF FFFFFFFF 00000000 00000000 00000000 00000000 00000000
	// 00000000 00000000 000000FF 000000FF FFFFFFFF FFFFFFFF FFFFFFFF FFFFFFFF 000000FF 00000000 00000000 00000000 00000000 00000000
	// 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000
	// 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000
	// 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000
}
