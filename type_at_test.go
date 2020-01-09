package font8x8_test

import (
	"github.com/reiver/go-byteliteral"
	"github.com/reiver/go-font8x8"

	"image/color"

	"testing"
)

func TestTypeAt(t *testing.T) {

	var black = color.Gray{0}
	var WHITE = color.Gray{255}

	tests := []struct{
		Font font8x8.Type
		Expected [8][8]color.Gray
	}{
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},









		{
			Font: font8x8.Type{
				byteliteral.B00000001,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{WHITE,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000010,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,WHITE,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000100,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,WHITE,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00001000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,WHITE,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00010000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,WHITE,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00100000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,WHITE,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B01000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,WHITE,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B10000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,WHITE},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},









		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000001,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{WHITE,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000010,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,WHITE,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000100,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,WHITE,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00001000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,WHITE,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00010000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,WHITE,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00100000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,WHITE,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B01000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,WHITE,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B10000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,WHITE},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},









		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000001,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{WHITE,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000010,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,WHITE,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000100,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,WHITE,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00001000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,WHITE,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00010000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,WHITE,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00100000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,WHITE,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B01000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,WHITE,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B10000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,WHITE},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},









		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000001,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{WHITE,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000010,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,WHITE,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000100,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,WHITE,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00001000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,WHITE,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00010000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,WHITE,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00100000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,WHITE,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B01000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,WHITE,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B10000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,WHITE},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},









		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000001,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{WHITE,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000010,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,WHITE,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000100,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,WHITE,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00001000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,WHITE,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00010000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,WHITE,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00100000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,WHITE,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B01000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,WHITE,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B10000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,WHITE},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},









		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000001,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{WHITE,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000010,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,WHITE,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000100,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,WHITE,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00001000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,WHITE,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00010000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,WHITE,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00100000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,WHITE,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B01000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,WHITE,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B10000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,WHITE},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},









		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000001,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{WHITE,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000010,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,WHITE,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000100,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,WHITE,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00001000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,WHITE,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00010000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,WHITE,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00100000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,WHITE,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B01000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,WHITE,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B10000000,
				byteliteral.B00000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,WHITE},
				[8]color.Gray{black,black,black,black,black,black,black,black},
			},
		},









		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000001,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{WHITE,black,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000010,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,WHITE,black,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000100,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,WHITE,black,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00001000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,WHITE,black,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00010000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,WHITE,black,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00100000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,WHITE,black,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B01000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,WHITE,black},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B00000000,
				byteliteral.B10000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,black,black,WHITE},
			},
		},









		{
			Font: font8x8.Type{
				byteliteral.B00000001,
				byteliteral.B00000011,
				byteliteral.B00000111,
				byteliteral.B00001111,
				byteliteral.B00011111,
				byteliteral.B00111111,
				byteliteral.B01111111,
				byteliteral.B11111111,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{WHITE,black,black,black,black,black,black,black},
				[8]color.Gray{WHITE,WHITE,black,black,black,black,black,black},
				[8]color.Gray{WHITE,WHITE,WHITE,black,black,black,black,black},
				[8]color.Gray{WHITE,WHITE,WHITE,WHITE,black,black,black,black},
				[8]color.Gray{WHITE,WHITE,WHITE,WHITE,WHITE,black,black,black},
				[8]color.Gray{WHITE,WHITE,WHITE,WHITE,WHITE,WHITE,black,black},
				[8]color.Gray{WHITE,WHITE,WHITE,WHITE,WHITE,WHITE,WHITE,black},
				[8]color.Gray{WHITE,WHITE,WHITE,WHITE,WHITE,WHITE,WHITE,WHITE},
			},
		},
		{
			Font: font8x8.Type{
				byteliteral.B11111111,
				byteliteral.B11111110,
				byteliteral.B11111100,
				byteliteral.B11111000,
				byteliteral.B11110000,
				byteliteral.B11100000,
				byteliteral.B11000000,
				byteliteral.B10000000,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{WHITE,WHITE,WHITE,WHITE,WHITE,WHITE,WHITE,WHITE},
				[8]color.Gray{black,WHITE,WHITE,WHITE,WHITE,WHITE,WHITE,WHITE},
				[8]color.Gray{black,black,WHITE,WHITE,WHITE,WHITE,WHITE,WHITE},
				[8]color.Gray{black,black,black,WHITE,WHITE,WHITE,WHITE,WHITE},
				[8]color.Gray{black,black,black,black,WHITE,WHITE,WHITE,WHITE},
				[8]color.Gray{black,black,black,black,black,WHITE,WHITE,WHITE},
				[8]color.Gray{black,black,black,black,black,black,WHITE,WHITE},
				[8]color.Gray{black,black,black,black,black,black,black,WHITE},
			},
		},









		{
			Font: font8x8.Type{
				byteliteral.B11011111,
				byteliteral.B11111111,
				byteliteral.B00000011,
				byteliteral.B11100000,
				byteliteral.B00000000,
				byteliteral.B00000001,
				byteliteral.B01010100,
				byteliteral.B10000010,
			},
			Expected: [8][8]color.Gray{
				[8]color.Gray{WHITE,WHITE,WHITE,WHITE,WHITE,black,WHITE,WHITE},
				[8]color.Gray{WHITE,WHITE,WHITE,WHITE,WHITE,WHITE,WHITE,WHITE},
				[8]color.Gray{WHITE,WHITE,black,black,black,black,black,black},
				[8]color.Gray{black,black,black,black,black,WHITE,WHITE,WHITE},
				[8]color.Gray{black,black,black,black,black,black,black,black},
				[8]color.Gray{WHITE,black,black,black,black,black,black,black},
				[8]color.Gray{black,black,WHITE,black,WHITE,black,WHITE,black},
				[8]color.Gray{black,black,black,black,black,black,black,WHITE},
			},
		},
	}

	for testNumber, test := range tests {

		for y:=0; y<7; y++ {
			for x:=0; x<7; x++ {

				actual := test.Font.At(x,y)
				if expected := test.Expected[y][x]; expected != actual  {
					t.Errorf("For test #%d, color actually at (x,y)=(%d,%d) was not what was expected.", testNumber, x, y)
					t.Logf("FONT:...\n01234567\n%z", test.Font)
					t.Logf("EXPECTED: %v", expected)
					t.Logf("ACTUAL:   %v", actual)
					continue
				}
			}
		}

	}
}
