package font8x8

import (
	"github.com/reiver/go-byteliteral"

	"fmt"
	"io"
	"strings"
)

func (receiver Type) Format(f fmt.State, c rune) {
	switch c {
	case 'z':
		var builder strings.Builder

		for _, b := range receiver {
			switch b & byteliteral.B00000001 {
			case 0:
				builder.WriteRune('░')
			default:
				builder.WriteRune('█')
			}
			switch b & byteliteral.B00000010 {
			case 0:
				builder.WriteRune('░')
			default:
				builder.WriteRune('█')
			}
			switch b & byteliteral.B00000100 {
			case 0:
				builder.WriteRune('░')
			default:
				builder.WriteRune('█')
			}
			switch b & byteliteral.B00001000 {
			case 0:
				builder.WriteRune('░')
			default:
				builder.WriteRune('█')
			}
			switch b & byteliteral.B00010000 {
			case 0:
				builder.WriteRune('░')
			default:
				builder.WriteRune('█')
			}
			switch b & byteliteral.B00100000 {
			case 0:
				builder.WriteRune('░')
			default:
				builder.WriteRune('█')
			}
			switch b & byteliteral.B01000000 {
			case 0:
				builder.WriteRune('░')
			default:
				builder.WriteRune('█')
			}
			switch b & byteliteral.B10000000 {
			case 0:
				builder.WriteRune('░')
			default:
				builder.WriteRune('█')
			}
			builder.WriteRune('\n')
		}

		io.WriteString(f, builder.String())
	default:
		fmt.Fprintf(f, "[8]byte{%x,%x,%x,%x,%x,%x,%x,%x}", receiver[0], receiver[1], receiver[2], receiver[3], receiver[4], receiver[5], receiver[6], receiver[7])
	}
}
