package font8x8

import (
	"image"
	"image/color"
)

const (
	minX = 0
	maxX = 8
	minY = 0
	maxY = 8
)

// At helps make font8x8.Type fit the image.Image interface.
//
// See: https://golang.org/pkg/image/#Image
func (receiver Type) At(x, y int) color.Color {
	const black = 0
	const white = 255

	if x < minX || maxX < x || y < minY || maxY < y {
		return color.Gray{black}
	}

	result := receiver[y] & (0x01 << byte(x))

	if 0 == result {
		return color.Gray{black}
	}

	return color.Gray{white}
}

// Bounds helps make font8x8.Type fit the image.Image interface.
//
// See: https://golang.org/pkg/image/#Image
func (receiver Type) Bounds() image.Rectangle {
	return image.Rect(minX,minY, maxX,maxY)
}

func (receiver Type) ColorModel() color.Model {
	return color.GrayModel
}
