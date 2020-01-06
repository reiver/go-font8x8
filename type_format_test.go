package font8x8_test

import (
	"github.com/reiver/go-byteliteral"
	"github.com/reiver/go-font8x8"

	"fmt"

	"testing"
)

func TestTypeFormatA(t *testing.T) {

	tests := []struct {
		Font font8x8.Type
		Expected string
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
			Expected:
				`░░░░░░░░`+"\n"+
				`░░░░░░░░`+"\n"+
				`░░░░░░░░`+"\n"+
				`░░░░░░░░`+"\n"+
				`░░░░░░░░`+"\n"+
				`░░░░░░░░`+"\n"+
				`░░░░░░░░`+"\n"+
				`░░░░░░░░`+"\n",
		},



		{
			Font: font8x8.Type{
				byteliteral.B10000000,
				byteliteral.B01000000,
				byteliteral.B00100000,
				byteliteral.B00010000,
				byteliteral.B00001000,
				byteliteral.B00000100,
				byteliteral.B00000010,
				byteliteral.B00000001,
			},
			Expected:
				`█░░░░░░░`+"\n"+
				`░█░░░░░░`+"\n"+
				`░░█░░░░░`+"\n"+
				`░░░█░░░░`+"\n"+
				`░░░░█░░░`+"\n"+
				`░░░░░█░░`+"\n"+
				`░░░░░░█░`+"\n"+
				`░░░░░░░█`+"\n",
		},



		{
			Font: font8x8.Type{
				byteliteral.B00000001,
				byteliteral.B00000010,
				byteliteral.B00000100,
				byteliteral.B00001000,
				byteliteral.B00010000,
				byteliteral.B00100000,
				byteliteral.B01000000,
				byteliteral.B10000000,
			},
			Expected:
				`░░░░░░░█`+"\n"+
				`░░░░░░█░`+"\n"+
				`░░░░░█░░`+"\n"+
				`░░░░█░░░`+"\n"+
				`░░░█░░░░`+"\n"+
				`░░█░░░░░`+"\n"+
				`░█░░░░░░`+"\n"+
				`█░░░░░░░`+"\n",
		},



		{
			Font: font8x8.Type{
				byteliteral.B10000001,
				byteliteral.B01000010,
				byteliteral.B00100100,
				byteliteral.B00011000,
				byteliteral.B00011000,
				byteliteral.B00100100,
				byteliteral.B01000010,
				byteliteral.B10000001,
			},
			Expected:
				`█░░░░░░█`+"\n"+
				`░█░░░░█░`+"\n"+
				`░░█░░█░░`+"\n"+
				`░░░██░░░`+"\n"+
				`░░░██░░░`+"\n"+
				`░░█░░█░░`+"\n"+
				`░█░░░░█░`+"\n"+
				`█░░░░░░█`+"\n",
		},



		{
			Font: font8x8.Type{
				byteliteral.B11111111,
				byteliteral.B11111111,
				byteliteral.B11111111,
				byteliteral.B11111111,
				byteliteral.B11111111,
				byteliteral.B11111111,
				byteliteral.B11111111,
				byteliteral.B11111111,
			},
			Expected:
				`████████`+"\n"+
				`████████`+"\n"+
				`████████`+"\n"+
				`████████`+"\n"+
				`████████`+"\n"+
				`████████`+"\n"+
				`████████`+"\n"+
				`████████`+"\n",
		},



		{
			Font: font8x8.Type{
				byteliteral.B10000000,
				byteliteral.B10000000,
				byteliteral.B10000000,
				byteliteral.B10000000,
				byteliteral.B10000000,
				byteliteral.B10000000,
				byteliteral.B10000000,
				byteliteral.B10000000,
			},
			Expected:
				`█░░░░░░░`+"\n"+
				`█░░░░░░░`+"\n"+
				`█░░░░░░░`+"\n"+
				`█░░░░░░░`+"\n"+
				`█░░░░░░░`+"\n"+
				`█░░░░░░░`+"\n"+
				`█░░░░░░░`+"\n"+
				`█░░░░░░░`+"\n",
		},
		{
			Font: font8x8.Type{
				byteliteral.B01000000,
				byteliteral.B01000000,
				byteliteral.B01000000,
				byteliteral.B01000000,
				byteliteral.B01000000,
				byteliteral.B01000000,
				byteliteral.B01000000,
				byteliteral.B01000000,
			},
			Expected:
				`░█░░░░░░`+"\n"+
				`░█░░░░░░`+"\n"+
				`░█░░░░░░`+"\n"+
				`░█░░░░░░`+"\n"+
				`░█░░░░░░`+"\n"+
				`░█░░░░░░`+"\n"+
				`░█░░░░░░`+"\n"+
				`░█░░░░░░`+"\n",
		},
		{
			Font: font8x8.Type{
				byteliteral.B00100000,
				byteliteral.B00100000,
				byteliteral.B00100000,
				byteliteral.B00100000,
				byteliteral.B00100000,
				byteliteral.B00100000,
				byteliteral.B00100000,
				byteliteral.B00100000,
			},
			Expected:
				`░░█░░░░░`+"\n"+
				`░░█░░░░░`+"\n"+
				`░░█░░░░░`+"\n"+
				`░░█░░░░░`+"\n"+
				`░░█░░░░░`+"\n"+
				`░░█░░░░░`+"\n"+
				`░░█░░░░░`+"\n"+
				`░░█░░░░░`+"\n",
		},
		{
			Font: font8x8.Type{
				byteliteral.B00010000,
				byteliteral.B00010000,
				byteliteral.B00010000,
				byteliteral.B00010000,
				byteliteral.B00010000,
				byteliteral.B00010000,
				byteliteral.B00010000,
				byteliteral.B00010000,
			},
			Expected:
				`░░░█░░░░`+"\n"+
				`░░░█░░░░`+"\n"+
				`░░░█░░░░`+"\n"+
				`░░░█░░░░`+"\n"+
				`░░░█░░░░`+"\n"+
				`░░░█░░░░`+"\n"+
				`░░░█░░░░`+"\n"+
				`░░░█░░░░`+"\n",
		},
		{
			Font: font8x8.Type{
				byteliteral.B00001000,
				byteliteral.B00001000,
				byteliteral.B00001000,
				byteliteral.B00001000,
				byteliteral.B00001000,
				byteliteral.B00001000,
				byteliteral.B00001000,
				byteliteral.B00001000,
			},
			Expected:
				`░░░░█░░░`+"\n"+
				`░░░░█░░░`+"\n"+
				`░░░░█░░░`+"\n"+
				`░░░░█░░░`+"\n"+
				`░░░░█░░░`+"\n"+
				`░░░░█░░░`+"\n"+
				`░░░░█░░░`+"\n"+
				`░░░░█░░░`+"\n",
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000100,
				byteliteral.B00000100,
				byteliteral.B00000100,
				byteliteral.B00000100,
				byteliteral.B00000100,
				byteliteral.B00000100,
				byteliteral.B00000100,
				byteliteral.B00000100,
			},
			Expected:
				`░░░░░█░░`+"\n"+
				`░░░░░█░░`+"\n"+
				`░░░░░█░░`+"\n"+
				`░░░░░█░░`+"\n"+
				`░░░░░█░░`+"\n"+
				`░░░░░█░░`+"\n"+
				`░░░░░█░░`+"\n"+
				`░░░░░█░░`+"\n",
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000010,
				byteliteral.B00000010,
				byteliteral.B00000010,
				byteliteral.B00000010,
				byteliteral.B00000010,
				byteliteral.B00000010,
				byteliteral.B00000010,
				byteliteral.B00000010,
			},
			Expected:
				`░░░░░░█░`+"\n"+
				`░░░░░░█░`+"\n"+
				`░░░░░░█░`+"\n"+
				`░░░░░░█░`+"\n"+
				`░░░░░░█░`+"\n"+
				`░░░░░░█░`+"\n"+
				`░░░░░░█░`+"\n"+
				`░░░░░░█░`+"\n",
		},
		{
			Font: font8x8.Type{
				byteliteral.B00000001,
				byteliteral.B00000001,
				byteliteral.B00000001,
				byteliteral.B00000001,
				byteliteral.B00000001,
				byteliteral.B00000001,
				byteliteral.B00000001,
				byteliteral.B00000001,
			},
			Expected:
				`░░░░░░░█`+"\n"+
				`░░░░░░░█`+"\n"+
				`░░░░░░░█`+"\n"+
				`░░░░░░░█`+"\n"+
				`░░░░░░░█`+"\n"+
				`░░░░░░░█`+"\n"+
				`░░░░░░░█`+"\n"+
				`░░░░░░░█`+"\n",
		},
	}

	for testNumber, test := range tests {

		actual := fmt.Sprintf("%z", test.Font)

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, did not actually get expected output.", testNumber)
			t.Log("EXPECTED:")
			t.Logf("\n%s", expected)
			t.Log("ACTUAL:")
			t.Logf("\n%s", actual)
			continue
		}
	}
}
