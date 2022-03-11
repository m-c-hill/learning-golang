# Chapter 1: Setting Up Your Go Environment

## 2.0 Built-in Types

Types can be viewed as value templates, and values can be viewed as type instances.

### 2.1 The Zero Value

Variable declared but not assigned a value is given an explicit "zero value". Each type has a zero value:
- boolean type is false.
- numeric type is zero, though zeros of different numeric types may have different sizes in memory.
- string type is an empty string.


### 2.2 Literals
A literal of a value is a text representation of the value in code:
- Integer literals: sequence of numbers, normally base 10. Can be prefixed with 0b for binary, 0o for octal and 0x for hexadecimal.
- Floating point literals: decimal point to indicate fractional portion of the value. Can also include exponent e (eg. 6.03e23)
- Rune literals: single characters surrounded by single quotes. Can include unicode characters, oct nums, hex nums, unicode nums.
- Raw string litetal: delimited with backquotes (\`) and can contain any character (except a backquote). Can be multiline.

All literals have a default type:
- The default type of a string literal is `string`.
- The default type of a boolean literal is `bool`.
- The default type of an integer literal is `int`.
- The default type of a rune literal is `rune` (a.k.a., `int32`).
- The default type of a floating-point literal is `float64`.

#### 2.3 Booleans
The `bool` type represents boolean variables, which are `true` or `false`.


#### 2.4 Integer Types
There are 11 built-in integer numeric types: `int8`, `uint8`, `int16`, `uint16`, `int32`, `uint32`, `int64`, `uint64`, `int`, `uint` and `uintptr`.

If a value occupies N bits in memory, we say the size of the value is N bits. The sizes of all values of a type are always the same, so value sizes are often called as type sizes.

The primitive data types prefixed with "u" are unsigned versions with the same bit sizes. Effectively, this means they cannot store negative numbers, but on the other hand they can store positive numbers twice as large as their signed counterparts. For example, int8 hase range (-128, 127) while uint8 has range (0, 255).

Note:
- `byte` is a built-in alias of `uint8`. We can view `byte` and `uint8` as the same type.
- `rune` is a built-in alias of `int32`. We can view `rune` and `int32` as the same type.

The standard operators work on numeric types +, -, \*, / and %, along with comparisons ==, !=, <, <=, > and >=. Note: due to floating point errors, == should not be used on floating point types (test within a range instead).


#### 2.5 Floating Point Types
There are two floating point types: `float32` and `float64`.

#### 2.6 Escapes in Literals
The following characters may be combined with the escape character `\` to give the following special characters:
```
\a   (Unicode value 0x07) alert or bell
\b   (Unicode value 0x08) backspace
\f   (Unicode value 0x0C) form feed
\n   (Unicode value 0x0A) line feed or newline
\r   (Unicode value 0x0D) carriage return
\t   (Unicode value 0x09) horizontal tab
\v   (Unicode value 0x0b) vertical tab
\\   (Unicode value 0x5c) backslash
\'   (Unicode value 0x27) single quote
```


### 2.7 Type Conversion
Go does not allow _automatic type promotion_ between variables (where numeric types are automatically converted from one to another when needed). Type conversion must be explicit using `<type>(<variable>)`:
```
var x int = 10  // x = 10
var y float64 = 30.2  // y = 30.2
var z float64 = float64(x) + y  // z = 40.2
var d int = x + int(y)  // No automatic conversion of numerics, running it will print "40.2 40"
```

### 2.8 var Versus :=
Variables can be declared 'verbosely':
`var x int = 10`

They can also be interpretted. Here, the defalt type of an integer literal is `int`, so x is of type `int`:
`var x = 10`

A variable can also be declared without assigning a value (automatically assigned the zero value)
`var x int`

Multiline variable declaration:
```
var a, b int = 10, 20
var x, y = 10, "hello"
```

Short declaration format:
`x := 0`

It is not legal to use `:=` outside of a function.

### 2.9 Const
To declare a value as immutable, use the `const` keyword.

Limitations: constants in Go are a way to give names to literals, they can only hold values that the compiler can *can figure out at compile time*. Therefore, they can only be assigned:
- Numeric literals
- `true` and `false`
- Strings
- Runes
- Built-in functions such as `complex`, `real`, `imag`, `len` and `cap`

```
const (
    idKey = "id"
    nameKey = "name"
)

const x = 20 * 10
```

### 2.10 Types and Untyped Constants
Constants can be _types_ or _untyped_. Untyped contants work like literals - they have no type of their own by do have a default type when no other type can be inferred.

```
const x = 10

var y int = x
var z flaot64 = x
var d byte = x  // byte is an alias of uint8
```

A typed constant can only be assigned to other variables of that type:
```
const typedX int = 10
```

## References
https://go101.org/article/basic-types-and-value-literals.html