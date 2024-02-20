package font8x8

import (
	"fmt"
	"io"
	"strings"
)

func (receiver Type) Format(f fmt.State, c rune) {
	switch c {
	case 'z':
		var builder strings.Builder

		for _, b := range receiver {
			switch b & 0b00000001 {
			case 0:
				builder.WriteRune('░')
			default:
				builder.WriteRune('█')
			}
			switch b & 0b00000010 {
			case 0:
				builder.WriteRune('░')
			default:
				builder.WriteRune('█')
			}
			switch b & 0b00000100 {
			case 0:
				builder.WriteRune('░')
			default:
				builder.WriteRune('█')
			}
			switch b & 0b00001000 {
			case 0:
				builder.WriteRune('░')
			default:
				builder.WriteRune('█')
			}
			switch b & 0b00010000 {
			case 0:
				builder.WriteRune('░')
			default:
				builder.WriteRune('█')
			}
			switch b & 0b00100000 {
			case 0:
				builder.WriteRune('░')
			default:
				builder.WriteRune('█')
			}
			switch b & 0b01000000 {
			case 0:
				builder.WriteRune('░')
			default:
				builder.WriteRune('█')
			}
			switch b & 0b10000000 {
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
