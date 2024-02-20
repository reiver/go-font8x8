package font8x8_test

import (
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b00000001,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b00000010,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b00000100,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b00001000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b00010000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b00100000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b01000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b10000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b00000001,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b00000010,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b00000100,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b00001000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b00010000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b00100000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b01000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b10000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b00000001,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b00000010,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b00000100,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b00001000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b00010000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b00100000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b01000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b10000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000001,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000010,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000100,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b00001000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b00010000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b00100000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b01000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b10000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000001,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000010,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000100,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00001000,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00010000,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00100000,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b01000000,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b10000000,
				0b00000000,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000001,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000010,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000100,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00001000,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00010000,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00100000,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b01000000,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b10000000,
				0b00000000,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000001,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000010,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000100,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00001000,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00010000,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00100000,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b01000000,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b10000000,
				0b00000000,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000001,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000010,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000100,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00001000,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00010000,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00100000,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b01000000,
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
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b00000000,
				0b10000000,
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
				0b00000001,
				0b00000011,
				0b00000111,
				0b00001111,
				0b00011111,
				0b00111111,
				0b01111111,
				0b11111111,
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
				0b11111111,
				0b11111110,
				0b11111100,
				0b11111000,
				0b11110000,
				0b11100000,
				0b11000000,
				0b10000000,
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
				0b11011111,
				0b11111111,
				0b00000011,
				0b11100000,
				0b00000000,
				0b00000001,
				0b01010100,
				0b10000010,
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
