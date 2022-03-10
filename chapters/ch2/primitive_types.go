/* Primitive types in Go */


// ==============================
// Integer value literals
// ==============================

0xF // the hex form (starts with a "0x" or "0X")
0XF
017 // the octal form (starts with a "0", "0o" or "0O")
0o17
0O17
0b1111 // the binary form (starts with a "0b" or "0B")
0B1111
15 // the decimal form (starts without a "0")


// ==============================
// Floating point value literals
// ==============================

1.23
01.23 // == 1.23
.23
1.
// An "e" or "E" starts the exponent part (10-based).
1.23e2  // == 123.0
123E2   // == 12300.0
123.E+2 // == 12300.0
1e-1    // == 0.1
.1e0    // == 0.1
0010e-2 // == 0.1
0e+5    // == 0.0


// ==============================
// Rune value literals
// ==============================

'a' // an English character
'π' // a mathematical character
'众' // a Chinese character

// 141 is the octal representation of decimal number 97.
'\141'
// 61 is the hex representation of decimal number 97.
'\x61'
'\u0061'
'\U00000061'


// ==============================
// String value literals
// ==============================

// The interpreted form.
"Hello\nworld!\n\"你好世界\""

// The raw form.
`Hello
world!
"你好世界"`

// The following interpreted string literals are equivalent.
"\141\142\143"
"\x61\x62\x63"
"\x61b\x63"
"abc"


// ==============================
// Type conversion
// ==============================

var x int = 10  // x = 10
var y float64 = 30.2  // y = 30.2
var z float64 = float64(x) + y  // z = 40.2
var d int = x + int(y)  // No automatic conversion of numerics, running it will print "40.2 40"
