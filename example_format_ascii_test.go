package font8x8_test

import (
	"github.com/reiver/go-font8x8"

	"fmt"
)

func ExampleType_Format_ascii() {

	fmt.Print("U+0021 (!)   EXPLAMATION MARK\n\n")
	fmt.Printf("%z\n", font8x8.U0021)

	fmt.Print("U+0022 (\")   QUOTATION MARK\n\n")
	fmt.Printf("%z\n", font8x8.U0022)

	fmt.Print("U+0023 (#)   NUMBER SIGN\n\n")
	fmt.Printf("%z\n", font8x8.U0023)

	fmt.Print("U+0024 ($)   DOLLAR SIGN\n\n")
	fmt.Printf("%z\n", font8x8.U0024)

	fmt.Print("U+0025 (%)   PERCENT SIGN\n\n")
	fmt.Printf("%z\n", font8x8.U0025)

	fmt.Print("U+0026 (&)   AMPERSAND\n\n")
	fmt.Printf("%z\n", font8x8.U0026)

	fmt.Print("U+0027 (')   APOSTROPHE\n\n")
	fmt.Printf("%z\n", font8x8.U0027)

	fmt.Print("U+0028 (()   LEFT PARENTHESIS\n\n")
	fmt.Printf("%z\n", font8x8.U0028)

	fmt.Print("U+0029 ())   RIGHT PARENTHESIS\n\n")
	fmt.Printf("%z\n", font8x8.U0029)

	fmt.Print("U+002A (*)   ASTERISK\n\n")
	fmt.Printf("%z\n", font8x8.U002A)

	fmt.Print("U+002B (+)   PLUS SIGN\n\n")
	fmt.Printf("%z\n", font8x8.U002B)

	fmt.Print("U+002C (,)   COMMA\n\n")
	fmt.Printf("%z\n", font8x8.U002C)

	fmt.Print("U+002D (-)   HYPHEN-MINUS\n\n")
	fmt.Printf("%z\n", font8x8.U002D)

	fmt.Print("U+002E (.)   FULL STOP\n\n")
	fmt.Printf("%z\n", font8x8.U002E)

	fmt.Print("U+002F (/)   SOLIDUS\n\n")
	fmt.Printf("%z\n", font8x8.U002F)

	fmt.Print("U+0030 (0)   ARABIC DIGIT ZERO\n\n")
	fmt.Printf("%z\n", font8x8.U0030)

	fmt.Print("U+0031 (1)   ARABIC DIGIT ONE\n\n")
	fmt.Printf("%z\n", font8x8.U0031)

	fmt.Print("U+0032 (2)   ARABIC DIGIT TWO\n\n")
	fmt.Printf("%z\n", font8x8.U0032)

	fmt.Print("U+0033 (3)   ARABIC DIGIT THREE\n\n")
	fmt.Printf("%z\n", font8x8.U0033)

	fmt.Print("U+0034 (4)   ARABIC DIGIT FOUR\n\n")
	fmt.Printf("%z\n", font8x8.U0034)

	fmt.Print("U+0035 (5)   ARABIC DIGIT FIVE\n\n")
	fmt.Printf("%z\n", font8x8.U0035)

	fmt.Print("U+0036 (6)   ARABIC DIGIT SIX\n\n")
	fmt.Printf("%z\n", font8x8.U0036)

	fmt.Print("U+0037 (7)   ARABIC DIGIT SEVEN\n\n")
	fmt.Printf("%z\n", font8x8.U0037)

	fmt.Print("U+0038 (8)   ARABIC DIGIT EIGHT\n\n")
	fmt.Printf("%z\n", font8x8.U0038)

	fmt.Print("U+0039 (9)   ARABIC DIGIT NINE\n\n")
	fmt.Printf("%z\n", font8x8.U0039)

	// Output:
	//
	// U+0021 (!)   EXPLAMATION MARK
	//
	// ░░░██░░░
	// ░░████░░
	// ░░████░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░░░░░░
	// ░░░██░░░
	// ░░░░░░░░
	//
	// U+0022 (")   QUOTATION MARK
	//
	// ░██░██░░
	// ░██░██░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+0023 (#)   NUMBER SIGN
	//
	// ░██░██░░
	// ░██░██░░
	// ███████░
	// ░██░██░░
	// ███████░
	// ░██░██░░
	// ░██░██░░
	// ░░░░░░░░
	//
	// U+0024 ($)   DOLLAR SIGN
	//
	// ░░██░░░░
	// ░█████░░
	// ██░░░░░░
	// ░████░░░
	// ░░░░██░░
	// █████░░░
	// ░░██░░░░
	// ░░░░░░░░
	//
	// U+0025 (%)   PERCENT SIGN
	//
	// ░░░░░░░░
	// ██░░░██░
	// ██░░██░░
	// ░░░██░░░
	// ░░██░░░░
	// ░██░░██░
	// ██░░░██░
	// ░░░░░░░░
	//
	// U+0026 (&)   AMPERSAND
	//
	// ░░███░░░
	// ░██░██░░
	// ░░███░░░
	// ░███░██░
	// ██░███░░
	// ██░░██░░
	// ░███░██░
	// ░░░░░░░░
	//
	// U+0027 (')   APOSTROPHE
	//
	// ░██░░░░░
	// ░██░░░░░
	// ██░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+0028 (()   LEFT PARENTHESIS
	//
	// ░░░██░░░
	// ░░██░░░░
	// ░██░░░░░
	// ░██░░░░░
	// ░██░░░░░
	// ░░██░░░░
	// ░░░██░░░
	// ░░░░░░░░
	//
	// U+0029 ())   RIGHT PARENTHESIS
	//
	// ░██░░░░░
	// ░░██░░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░██░░░░
	// ░██░░░░░
	// ░░░░░░░░
	//
	// U+002A (*)   ASTERISK
	//
	// ░░░░░░░░
	// ░██░░██░
	// ░░████░░
	// ████████
	// ░░████░░
	// ░██░░██░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+002B (+)   PLUS SIGN
	//
	// ░░░░░░░░
	// ░░██░░░░
	// ░░██░░░░
	// ██████░░
	// ░░██░░░░
	// ░░██░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+002C (,)   COMMA
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░██░░░░
	// ░░██░░░░
	// ░██░░░░░
	//
	// U+002D (-)   HYPHEN-MINUS
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ██████░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+002E (.)   FULL STOP
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░██░░░░
	// ░░██░░░░
	// ░░░░░░░░
	//
	// U+002F (/)   SOLIDUS
	//
	// ░░░░░██░
	// ░░░░██░░
	// ░░░██░░░
	// ░░██░░░░
	// ░██░░░░░
	// ██░░░░░░
	// █░░░░░░░
	// ░░░░░░░░
	//
	// U+0030 (0)   ARABIC DIGIT ZERO
	//
	// ░█████░░
	// ██░░░██░
	// ██░░███░
	// ██░████░
	// ████░██░
	// ███░░██░
	// ░█████░░
	// ░░░░░░░░
	//
	// U+0031 (1)   ARABIC DIGIT ONE
	//
	// ░░██░░░░
	// ░███░░░░
	// ░░██░░░░
	// ░░██░░░░
	// ░░██░░░░
	// ░░██░░░░
	// ██████░░
	// ░░░░░░░░
	//
	// U+0032 (2)   ARABIC DIGIT TWO
	//
	// ░████░░░
	// ██░░██░░
	// ░░░░██░░
	// ░░███░░░
	// ░██░░░░░
	// ██░░██░░
	// ██████░░
	// ░░░░░░░░
	//
	// U+0033 (3)   ARABIC DIGIT THREE
	//
	// ░████░░░
	// ██░░██░░
	// ░░░░██░░
	// ░░███░░░
	// ░░░░██░░
	// ██░░██░░
	// ░████░░░
	// ░░░░░░░░
	//
	// U+0034 (4)   ARABIC DIGIT FOUR
	//
	// ░░░███░░
	// ░░████░░
	// ░██░██░░
	// ██░░██░░
	// ███████░
	// ░░░░██░░
	// ░░░████░
	// ░░░░░░░░
	//
	// U+0035 (5)   ARABIC DIGIT FIVE
	//
	// ██████░░
	// ██░░░░░░
	// █████░░░
	// ░░░░██░░
	// ░░░░██░░
	// ██░░██░░
	// ░████░░░
	// ░░░░░░░░
	//
	// U+0036 (6)   ARABIC DIGIT SIX
	//
	// ░░███░░░
	// ░██░░░░░
	// ██░░░░░░
	// █████░░░
	// ██░░██░░
	// ██░░██░░
	// ░████░░░
	// ░░░░░░░░
	//
	// U+0037 (7)   ARABIC DIGIT SEVEN
	//
	// ██████░░
	// ██░░██░░
	// ░░░░██░░
	// ░░░██░░░
	// ░░██░░░░
	// ░░██░░░░
	// ░░██░░░░
	// ░░░░░░░░
	//
	// U+0038 (8)   ARABIC DIGIT EIGHT
	//
	// ░████░░░
	// ██░░██░░
	// ██░░██░░
	// ░████░░░
	// ██░░██░░
	// ██░░██░░
	// ░████░░░
	// ░░░░░░░░
	//
	// U+0039 (9)   ARABIC DIGIT NINE
	//
	// ░████░░░
	// ██░░██░░
	// ██░░██░░
	// ░█████░░
	// ░░░░██░░
	// ░░░██░░░
	// ░███░░░░
	// ░░░░░░░░
}
