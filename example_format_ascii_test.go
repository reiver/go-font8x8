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

	fmt.Print("U+003A (:)   COLON\n\n")
	fmt.Printf("%z\n", font8x8.U003A)

	fmt.Print("U+003B (;)   SEMICOLON\n\n")
	fmt.Printf("%z\n", font8x8.U003B)

	fmt.Print("U+003C (<)   LESS-THAN SIGN\n\n")
	fmt.Printf("%z\n", font8x8.U003C)

	fmt.Print("U+003D (=)   EQUALS SIGN\n\n")
	fmt.Printf("%z\n", font8x8.U003D)

	fmt.Print("U+003E (>)   GREATER-THAN SIGN\n\n")
	fmt.Printf("%z\n", font8x8.U003E)

	fmt.Print("U+003F (?)   QUESTION MARK\n\n")
	fmt.Printf("%z\n", font8x8.U003F)

	fmt.Print("U+0040 (@)   AT SIGN\n\n")
	fmt.Printf("%z\n", font8x8.U0040)

	fmt.Print("U+0041 (A)   LATIN CAPITAL LETTER A\n\n")
	fmt.Printf("%z\n", font8x8.U0041)

	fmt.Print("U+0042 (B)   LATIN CAPITAL LETTER B\n\n")
	fmt.Printf("%z\n", font8x8.U0042)

	fmt.Print("U+0043 (C)   LATIN CAPITAL LETTER C\n\n")
	fmt.Printf("%z\n", font8x8.U0043)

	fmt.Print("U+0044 (D)   LATIN CAPITAL LETTER D\n\n")
	fmt.Printf("%z\n", font8x8.U0044)

	fmt.Print("U+0045 (E)   LATIN CAPITAL LETTER E\n\n")

	fmt.Print("U+0046 (F)   LATIN CAPITAL LETTER F\n\n")

	fmt.Print("U+0047 (G)   LATIN CAPITAL LETTER G\n\n")

	fmt.Print("U+0048 (H)   LATIN CAPITAL LETTER H\n\n")

	fmt.Print("U+0049 (I)   LATIN CAPITAL LETTER I\n\n")

	fmt.Print("U+004A (J)   LATIN CAPITAL LETTER J\n\n")

	fmt.Print("U+004B (K)   LATIN CAPITAL LETTER K\n\n")

	fmt.Print("U+004C (L)   LATIN CAPITAL LETTER L\n\n")

	fmt.Print("U+004D (M)   LATIN CAPITAL LETTER M\n\n")

	fmt.Print("U+004E (N)   LATIN CAPITAL LETTER N\n\n")

	fmt.Print("U+004F (O)   LATIN CAPITAL LETTER O\n\n")

	fmt.Print("U+0050 (P)   LATIN CAPITAL LETTER P\n\n")

	fmt.Print("U+0051 (Q)   LATIN CAPITAL LETTER Q\n\n")

	fmt.Print("U+0052 (R)   LATIN CAPITAL LETTER R\n\n")

	fmt.Print("U+0053 (S)   LATIN CAPITAL LETTER S\n\n")

	fmt.Print("U+0054 (T)   LATIN CAPITAL LETTER T\n\n")

	fmt.Print("U+0055 (U)   LATIN CAPITAL LETTER U\n\n")

	fmt.Print("U+0056 (V)   LATIN CAPITAL LETTER V\n\n")

	fmt.Print("U+0057 (W)   LATIN CAPITAL LETTER W\n\n")

	fmt.Print("U+0058 (X)   LATIN CAPITAL LETTER X\n\n")

	fmt.Print("U+0059 (Y)   LATIN CAPITAL LETTER Y\n\n")

	fmt.Print("U+005A (Z)   LATIN CAPITAL LETTER Z\n\n")

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
	//
	// U+003A (:)   COLON
	//
	// ░░░░░░░░
	// ░░██░░░░
	// ░░██░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░██░░░░
	// ░░██░░░░
	// ░░░░░░░░
	//
	// U+003B (;)   SEMICOLON
	//
	// ░░░░░░░░
	// ░░██░░░░
	// ░░██░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░██░░░░
	// ░░██░░░░
	// ░██░░░░░
	//
	// U+003C (<)   LESS-THAN SIGN
	//
	// ░░░██░░░
	// ░░██░░░░
	// ░██░░░░░
	// ██░░░░░░
	// ░██░░░░░
	// ░░██░░░░
	// ░░░██░░░
	// ░░░░░░░░
	//
	// U+003D (=)   EQUALS SIGN
	//
	// ░░░░░░░░
	// ░░░░░░░░
	// ██████░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ██████░░
	// ░░░░░░░░
	// ░░░░░░░░
	//
	// U+003E (>)   GREATER-THAN SIGN
	//
	// ░██░░░░░
	// ░░██░░░░
	// ░░░██░░░
	// ░░░░██░░
	// ░░░██░░░
	// ░░██░░░░
	// ░██░░░░░
	// ░░░░░░░░
	//
	// U+003F (?)   QUESTION MARK
	//
	// ░████░░░
	// ██░░██░░
	// ░░░░██░░
	// ░░░██░░░
	// ░░██░░░░
	// ░░░░░░░░
	// ░░██░░░░
	// ░░░░░░░░
	//
	// U+0040 (@)   AT SIGN
	//
	// ░█████░░
	// ██░░░██░
	// ██░████░
	// ██░████░
	// ██░████░
	// ██░░░░░░
	// ░████░░░
	// ░░░░░░░░
	//
	// U+0041 (A)   LATIN CAPITAL LETTER A
	//
	// ░░██░░░░
	// ░████░░░
	// ██░░██░░
	// ██░░██░░
	// ██████░░
	// ██░░██░░
 	// ██░░██░░
	// ░░░░░░░░
	//
	// U+0042 (B)   LATIN CAPITAL LETTER B
	//
	// ██████░░
	// ░██░░██░
	// ░██░░██░
	// ░█████░░
	// ░██░░██░
	// ░██░░██░
	// ██████░░
	// ░░░░░░░░
	//
	// U+0043 (C)   LATIN CAPITAL LETTER C
	//
	// ░░████░░
	// ░██░░██░
	// ██░░░░░░
	// ██░░░░░░
	// ██░░░░░░
	// ░██░░██░
	// ░░████░░
	// ░░░░░░░░
	//
	// U+0044 (D)   LATIN CAPITAL LETTER D
	//
	// █████░░░
	// ░██░██░░
	// ░██░░██░
	// ░██░░██░
	// ░██░░██░
	// ░██░██░░
	// █████░░░
	// ░░░░░░░░
	//
	// U+0045 (E)   LATIN CAPITAL LETTER E
	//
	// U+0046 (F)   LATIN CAPITAL LETTER F
	//
	// U+0047 (G)   LATIN CAPITAL LETTER G
	//
	// U+0048 (H)   LATIN CAPITAL LETTER H
	//
	// U+0049 (I)   LATIN CAPITAL LETTER I
	//
	// U+004A (J)   LATIN CAPITAL LETTER J
	//
	// U+004B (K)   LATIN CAPITAL LETTER K
	//
	// U+004C (L)   LATIN CAPITAL LETTER L
	//
	// U+004D (M)   LATIN CAPITAL LETTER M
	//
	// U+004E (N)   LATIN CAPITAL LETTER N
	//
	// U+004F (O)   LATIN CAPITAL LETTER O
	//
	// U+0050 (P)   LATIN CAPITAL LETTER P
	//
	// U+0051 (Q)   LATIN CAPITAL LETTER Q
	//
	// U+0052 (R)   LATIN CAPITAL LETTER R
	//
	// U+0053 (S)   LATIN CAPITAL LETTER S
	//
	// U+0054 (T)   LATIN CAPITAL LETTER T
	//
	// U+0055 (U)   LATIN CAPITAL LETTER U
	//
	// U+0056 (V)   LATIN CAPITAL LETTER V
	//
	// U+0057 (W)   LATIN CAPITAL LETTER W
	//
	// U+0058 (X)   LATIN CAPITAL LETTER X
	//
	// U+0059 (Y)   LATIN CAPITAL LETTER Y
	//
	// U+005A (Z)   LATIN CAPITAL LETTER Z
}
