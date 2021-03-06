package font8x8

var (
	U0000 = Type{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} // U+0000 (NUL) NULL                        ^@ (␀) \0
	U0001 = Type{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} // U+0001 (SOH) START OF HEADING            ^A (␁)
	U0002 = Type{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} // U+0002 (STX) START OF TEXT               ^B (␂)
	U0003 = Type{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} // U+0003 (ETX) END OF TEXT                 ^C (␃)
	U0004 = Type{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} // U+0004 (EOT) END OF TRANSMISSION         ^D (␄)
	U0005 = Type{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} // U+0005 (ENQ) ENQUIRY                     ^E (␅)
	U0006 = Type{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} // U+0006 (ACK) ACKNOWLEDGE                 ^F (␆)
	U0007 = Type{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} // U+0007 (BEL) BELL                        ^G (␇) \a
	U0008 = Type{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} // U+0008 (BS)  BACKSPACE                   ^H (␈) \b
	U0009 = Type{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} // U+0009 (HT)  HORIZONTAL TABULATION       ^I (␉) \t
	U000A = Type{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} // U+000A (LF)  LINE FEED                   ^J (␊) \n
	U000B = Type{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} // U+000B (VT)  VERTICAL TABULATION         ^K (␋) \v
	U000C = Type{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} // U+000C (FF)  FORM FEED                   ^L (␌) \f
	U000D = Type{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} // U+000D (CR)  CARRIAGE RETURN             ^M (␍) \r
	U000E = Type{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} // U+000E (SO)  SHIFT OUT                   ^N (␎)
	U000F = Type{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} // U+000F (SI)  SHIFT IN                    ^O (␏)
	U0010 = Type{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} // U+0010 (DLE) DATA LINK ESCAPE            ^P (␐)
	U0011 = Type{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} // U+0011 (DC1) DEVICE CONTROL ONE          ^Q (␑)
	U0012 = Type{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} // U+0012 (DC2) DEVICE CONTROL TWO          ^R (␒)
	U0013 = Type{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} // U+0013 (DC3) DEVICE CONTROL THREE        ^S (␓)
	U0014 = Type{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} // U+0014 (DC4) DEVICE CONTROL FOUR         ^T (␔)
	U0015 = Type{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} // U+0015 (NAK) NEGATIVE ACKNOWLEDGE        ^U (␕)
	U0016 = Type{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} // U+0016 (SYN) SYNCHRONOUS IDLE            ^V (␖)
	U0017 = Type{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} // U+0017 (ETB) END OF TRANSMISSION BLOCK   ^W (␗)
	U0018 = Type{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} // U+0018 (CAN) CANCEL                      ^X (␘)
	U0019 = Type{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} // U+0019 (EM)  END OF MEDIUM               ^Y (␙)
	U001A = Type{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} // U+001A (SUB) SUBSTITUTE                  ^Z (␚)
	U001B = Type{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} // U+001B (ESC) ESCAPE                      ^[ (␛) \e
	U001C = Type{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} // U+001C (FS)  FILE SEPARATOR              ^\ (␜)
	U001D = Type{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} // U+001D (GS)  GROUP SEPARATOR             ^] (␝)
	U001E = Type{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} // U+001E (RS)  RECORD SEPARATOR            ^^ (␞)
	U001F = Type{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} // U+001F (US)  UNICODE SEPARATOR           ^_ (␟)
	U0020 = Type{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} // U+0020 (SP)  SPACE                          (␠)
	U0021 = Type{0x18, 0x3C, 0x3C, 0x18, 0x18, 0x00, 0x18, 0x00} // U+0021 (!)   EXPLAMATION MARK
	U0022 = Type{0x36, 0x36, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} // U+0022 (")   QUOTATION MARK
	U0023 = Type{0x36, 0x36, 0x7F, 0x36, 0x7F, 0x36, 0x36, 0x00} // U+0023 (#)   NUMBER SIGN
	U0024 = Type{0x0C, 0x3E, 0x03, 0x1E, 0x30, 0x1F, 0x0C, 0x00} // U+0024 ($)   DOLLAR SIGN
	U0025 = Type{0x00, 0x63, 0x33, 0x18, 0x0C, 0x66, 0x63, 0x00} // U+0025 (%)   PERCENT SIGN
	U0026 = Type{0x1C, 0x36, 0x1C, 0x6E, 0x3B, 0x33, 0x6E, 0x00} // U+0026 (&)   AMPERSAND
	U0027 = Type{0x06, 0x06, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00} // U+0027 (')   APOSTROPHE
	U0028 = Type{0x18, 0x0C, 0x06, 0x06, 0x06, 0x0C, 0x18, 0x00} // U+0028 (()   LEFT PARENTHESIS
	U0029 = Type{0x06, 0x0C, 0x18, 0x18, 0x18, 0x0C, 0x06, 0x00} // U+0029 ())   RIGHT PARENTHESIS
	U002A = Type{0x00, 0x66, 0x3C, 0xFF, 0x3C, 0x66, 0x00, 0x00} // U+002A (*)   ASTERISK
	U002B = Type{0x00, 0x0C, 0x0C, 0x3F, 0x0C, 0x0C, 0x00, 0x00} // U+002B (+)   PLUS SIGN
	U002C = Type{0x00, 0x00, 0x00, 0x00, 0x00, 0x0C, 0x0C, 0x06} // U+002C (,)   COMMA
	U002D = Type{0x00, 0x00, 0x00, 0x3F, 0x00, 0x00, 0x00, 0x00} // U+002D (-)   HYPHEN-MINUS
	U002E = Type{0x00, 0x00, 0x00, 0x00, 0x00, 0x0C, 0x0C, 0x00} // U+002E (.)   FULL STOP
	U002F = Type{0x60, 0x30, 0x18, 0x0C, 0x06, 0x03, 0x01, 0x00} // U+002F (/)   SOLIDUS
	U0030 = Type{0x3E, 0x63, 0x73, 0x7B, 0x6F, 0x67, 0x3E, 0x00} // U+0030 (0)   ARABIC DIGIT ZERO
	U0031 = Type{0x0C, 0x0E, 0x0C, 0x0C, 0x0C, 0x0C, 0x3F, 0x00} // U+0031 (1)   ARABIC DIGIT ONE
	U0032 = Type{0x1E, 0x33, 0x30, 0x1C, 0x06, 0x33, 0x3F, 0x00} // U+0032 (2)   ARABIC DIGIT TWO
	U0033 = Type{0x1E, 0x33, 0x30, 0x1C, 0x30, 0x33, 0x1E, 0x00} // U+0033 (3)   ARABIC DIGIT THREE
	U0034 = Type{0x38, 0x3C, 0x36, 0x33, 0x7F, 0x30, 0x78, 0x00} // U+0034 (4)   ARABIC DIGIT FOUR
	U0035 = Type{0x3F, 0x03, 0x1F, 0x30, 0x30, 0x33, 0x1E, 0x00} // U+0035 (5)   ARABIC DIGIT FIVE
	U0036 = Type{0x1C, 0x06, 0x03, 0x1F, 0x33, 0x33, 0x1E, 0x00} // U+0036 (6)   ARABIC DIGIT SIX
	U0037 = Type{0x3F, 0x33, 0x30, 0x18, 0x0C, 0x0C, 0x0C, 0x00} // U+0037 (7)   ARABIC DIGIT SEVEN
	U0038 = Type{0x1E, 0x33, 0x33, 0x1E, 0x33, 0x33, 0x1E, 0x00} // U+0038 (8)   ARABIC DIGIT EIGHT
	U0039 = Type{0x1E, 0x33, 0x33, 0x3E, 0x30, 0x18, 0x0E, 0x00} // U+0039 (9)   ARABIC DIGIT NINE
	U003A = Type{0x00, 0x0C, 0x0C, 0x00, 0x00, 0x0C, 0x0C, 0x00} // U+003A (:)   COLON
	U003B = Type{0x00, 0x0C, 0x0C, 0x00, 0x00, 0x0C, 0x0C, 0x06} // U+003B (;)   SEMICOLON
	U003C = Type{0x18, 0x0C, 0x06, 0x03, 0x06, 0x0C, 0x18, 0x00} // U+003C (<)   LESS-THAN SIGN
	U003D = Type{0x00, 0x00, 0x3F, 0x00, 0x00, 0x3F, 0x00, 0x00} // U+003D (=)   EQUALS SIGN
	U003E = Type{0x06, 0x0C, 0x18, 0x30, 0x18, 0x0C, 0x06, 0x00} // U+003E (>)   GREATER-THAN SIGN
	U003F = Type{0x1E, 0x33, 0x30, 0x18, 0x0C, 0x00, 0x0C, 0x00} // U+003F (?)   QUESTION MARK
	U0040 = Type{0x3E, 0x63, 0x7B, 0x7B, 0x7B, 0x03, 0x1E, 0x00} // U+0040 (@)   AT SIGN
	U0041 = Type{0x0C, 0x1E, 0x33, 0x33, 0x3F, 0x33, 0x33, 0x00} // U+0041 (A)   LATIN CAPITAL LETTER A
	U0042 = Type{0x3F, 0x66, 0x66, 0x3E, 0x66, 0x66, 0x3F, 0x00} // U+0042 (B)   LATIN CAPITAL LETTER B
	U0043 = Type{0x3C, 0x66, 0x03, 0x03, 0x03, 0x66, 0x3C, 0x00} // U+0043 (C)   LATIN CAPITAL LETTER C
	U0044 = Type{0x1F, 0x36, 0x66, 0x66, 0x66, 0x36, 0x1F, 0x00} // U+0044 (D)   LATIN CAPITAL LETTER D
	U0045 = Type{0x7F, 0x46, 0x16, 0x1E, 0x16, 0x46, 0x7F, 0x00} // U+0045 (E)   LATIN CAPITAL LETTER E
	U0046 = Type{0x7F, 0x46, 0x16, 0x1E, 0x16, 0x06, 0x0F, 0x00} // U+0046 (F)   LATIN CAPITAL LETTER F
	U0047 = Type{0x3C, 0x66, 0x03, 0x03, 0x73, 0x66, 0x7C, 0x00} // U+0047 (G)   LATIN CAPITAL LETTER G
	U0048 = Type{0x33, 0x33, 0x33, 0x3F, 0x33, 0x33, 0x33, 0x00} // U+0048 (H)   LATIN CAPITAL LETTER H
	U0049 = Type{0x1E, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x1E, 0x00} // U+0049 (I)   LATIN CAPITAL LETTER I
	U004A = Type{0x78, 0x30, 0x30, 0x30, 0x33, 0x33, 0x1E, 0x00} // U+004A (J)   LATIN CAPITAL LETTER J
	U004B = Type{0x67, 0x66, 0x36, 0x1E, 0x36, 0x66, 0x67, 0x00} // U+004B (K)   LATIN CAPITAL LETTER K
	U004C = Type{0x0F, 0x06, 0x06, 0x06, 0x46, 0x66, 0x7F, 0x00} // U+004C (L)   LATIN CAPITAL LETTER L
	U004D = Type{0x63, 0x77, 0x7F, 0x7F, 0x6B, 0x63, 0x63, 0x00} // U+004D (M)   LATIN CAPITAL LETTER M
	U004E = Type{0x63, 0x67, 0x6F, 0x7B, 0x73, 0x63, 0x63, 0x00} // U+004E (N)   LATIN CAPITAL LETTER N
	U004F = Type{0x1C, 0x36, 0x63, 0x63, 0x63, 0x36, 0x1C, 0x00} // U+004F (O)   LATIN CAPITAL LETTER O
	U0050 = Type{0x3F, 0x66, 0x66, 0x3E, 0x06, 0x06, 0x0F, 0x00} // U+0050 (P)   LATIN CAPITAL LETTER P
	U0051 = Type{0x1E, 0x33, 0x33, 0x33, 0x3B, 0x1E, 0x38, 0x00} // U+0051 (Q)   LATIN CAPITAL LETTER Q
	U0052 = Type{0x3F, 0x66, 0x66, 0x3E, 0x36, 0x66, 0x67, 0x00} // U+0052 (R)   LATIN CAPITAL LETTER R
	U0053 = Type{0x1E, 0x33, 0x07, 0x0E, 0x38, 0x33, 0x1E, 0x00} // U+0053 (S)   LATIN CAPITAL LETTER S
	U0054 = Type{0x3F, 0x2D, 0x0C, 0x0C, 0x0C, 0x0C, 0x1E, 0x00} // U+0054 (T)   LATIN CAPITAL LETTER T
	U0055 = Type{0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x3F, 0x00} // U+0055 (U)   LATIN CAPITAL LETTER U
	U0056 = Type{0x33, 0x33, 0x33, 0x33, 0x33, 0x1E, 0x0C, 0x00} // U+0056 (V)   LATIN CAPITAL LETTER V
	U0057 = Type{0x63, 0x63, 0x63, 0x6B, 0x7F, 0x77, 0x63, 0x00} // U+0057 (W)   LATIN CAPITAL LETTER W
	U0058 = Type{0x63, 0x63, 0x36, 0x1C, 0x1C, 0x36, 0x63, 0x00} // U+0058 (X)   LATIN CAPITAL LETTER X
	U0059 = Type{0x33, 0x33, 0x33, 0x1E, 0x0C, 0x0C, 0x1E, 0x00} // U+0059 (Y)   LATIN CAPITAL LETTER Y
	U005A = Type{0x7F, 0x63, 0x31, 0x18, 0x4C, 0x66, 0x7F, 0x00} // U+005A (Z)   LATIN CAPITAL LETTER Z
	U005B = Type{0x1E, 0x06, 0x06, 0x06, 0x06, 0x06, 0x1E, 0x00} // U+005B ([)   LEFT SQUARE BRACKET
	U005C = Type{0x03, 0x06, 0x0C, 0x18, 0x30, 0x60, 0x40, 0x00} // U+005C (\)   REVERSE SOLIDUS
	U005D = Type{0x1E, 0x18, 0x18, 0x18, 0x18, 0x18, 0x1E, 0x00} // U+005D (])   RIGHT SQUARE BRACKET
	U005E = Type{0x08, 0x1C, 0x36, 0x63, 0x00, 0x00, 0x00, 0x00} // U+005E (^)   CIRCUMFLEX ACCENT
	U005F = Type{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF} // U+005F (_)   LOW LINE
	U0060 = Type{0x0C, 0x0C, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00} // U+0060 (`)   GRAVE ACCENT
	U0061 = Type{0x00, 0x00, 0x1E, 0x30, 0x3E, 0x33, 0x6E, 0x00} // U+0061 (a)   LATIN SMALL LETTER a
	U0062 = Type{0x07, 0x06, 0x06, 0x3E, 0x66, 0x66, 0x3B, 0x00} // U+0062 (b)   LATIN SMALL LETTER b
	U0063 = Type{0x00, 0x00, 0x1E, 0x33, 0x03, 0x33, 0x1E, 0x00} // U+0063 (c)   LATIN SMALL LETTER c
	U0064 = Type{0x38, 0x30, 0x30, 0x3e, 0x33, 0x33, 0x6E, 0x00} // U+0064 (d)   LATIN SMALL LETTER d
	U0065 = Type{0x00, 0x00, 0x1E, 0x33, 0x3f, 0x03, 0x1E, 0x00} // U+0065 (e)   LATIN SMALL LETTER e
	U0066 = Type{0x1C, 0x36, 0x06, 0x0f, 0x06, 0x06, 0x0F, 0x00} // U+0066 (f)   LATIN SMALL LETTER f
	U0067 = Type{0x00, 0x00, 0x6E, 0x33, 0x33, 0x3E, 0x30, 0x1F} // U+0067 (g)   LATIN SMALL LETTER g
	U0068 = Type{0x07, 0x06, 0x36, 0x6E, 0x66, 0x66, 0x67, 0x00} // U+0068 (h)   LATIN SMALL LETTER h
	U0069 = Type{0x0C, 0x00, 0x0E, 0x0C, 0x0C, 0x0C, 0x1E, 0x00} // U+0069 (i)   LATIN SMALL LETTER i
	U006A = Type{0x30, 0x00, 0x30, 0x30, 0x30, 0x33, 0x33, 0x1E} // U+006A (j)   LATIN SMALL LETTER j
	U006B = Type{0x07, 0x06, 0x66, 0x36, 0x1E, 0x36, 0x67, 0x00} // U+006B (k)   LATIN SMALL LETTER k
	U006C = Type{0x0E, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x1E, 0x00} // U+006C (l)   LATIN SMALL LETTER l
	U006D = Type{0x00, 0x00, 0x33, 0x7F, 0x7F, 0x6B, 0x63, 0x00} // U+006D (m)   LATIN SMALL LETTER m
	U006E = Type{0x00, 0x00, 0x1F, 0x33, 0x33, 0x33, 0x33, 0x00} // U+006E (n)   LATIN SMALL LETTER n
	U006F = Type{0x00, 0x00, 0x1E, 0x33, 0x33, 0x33, 0x1E, 0x00} // U+006F (o)   LATIN SMALL LETTER o
	U0070 = Type{0x00, 0x00, 0x3B, 0x66, 0x66, 0x3E, 0x06, 0x0F} // U+0070 (p)   LATIN SMALL LETTER p
	U0071 = Type{0x00, 0x00, 0x6E, 0x33, 0x33, 0x3E, 0x30, 0x78} // U+0071 (q)   LATIN SMALL LETTER q
	U0072 = Type{0x00, 0x00, 0x3B, 0x6E, 0x66, 0x06, 0x0F, 0x00} // U+0072 (r)   LATIN SMALL LETTER r
	U0073 = Type{0x00, 0x00, 0x3E, 0x03, 0x1E, 0x30, 0x1F, 0x00} // U+0073 (s)   LATIN SMALL LETTER s
	U0074 = Type{0x08, 0x0C, 0x3E, 0x0C, 0x0C, 0x2C, 0x18, 0x00} // U+0074 (t)   LATIN SMALL LETTER t
	U0075 = Type{0x00, 0x00, 0x33, 0x33, 0x33, 0x33, 0x6E, 0x00} // U+0075 (u)   LATIN SMALL LETTER u
	U0076 = Type{0x00, 0x00, 0x33, 0x33, 0x33, 0x1E, 0x0C, 0x00} // U+0076 (v)   LATIN SMALL LETTER v
	U0077 = Type{0x00, 0x00, 0x63, 0x6B, 0x7F, 0x7F, 0x36, 0x00} // U+0077 (w)   LATIN SMALL LETTER w
	U0078 = Type{0x00, 0x00, 0x63, 0x36, 0x1C, 0x36, 0x63, 0x00} // U+0078 (x)   LATIN SMALL LETTER x
	U0079 = Type{0x00, 0x00, 0x33, 0x33, 0x33, 0x3E, 0x30, 0x1F} // U+0079 (y)   LATIN SMALL LETTER y
	U007A = Type{0x00, 0x00, 0x3F, 0x19, 0x0C, 0x26, 0x3F, 0x00} // U+007A (z)   LATIN SMALL LETTER z
	U007B = Type{0x38, 0x0C, 0x0C, 0x07, 0x0C, 0x0C, 0x38, 0x00} // U+007B ({)   LEFT CURLY BRACKET
	U007C = Type{0x18, 0x18, 0x18, 0x00, 0x18, 0x18, 0x18, 0x00} // U+007C (|)   VERTICAL LINE
	U007D = Type{0x07, 0x0C, 0x0C, 0x38, 0x0C, 0x0C, 0x07, 0x00} // U+007D (})   RIGHT CURLY BRACKET
	U007E = Type{0x6E, 0x3B, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} // U+007E (~)   TILDE
	U007F = Type{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} // U+007F (DEL) DELETE                      ^? (␡)
)
