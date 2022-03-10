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

#### 2.2.1 Booleans
The `bool` type represents boolean variables, which are `true` or `false`.


#### 2.2.2 Integer Types
There are 11 built-in integer numeric types: `int8`, `uint8`, `int16`, `uint16`, `int32`, `uint32`, `int64`, `uint64`, `int`, `uint` and `uintptr`.

If a value occupies N bits in memory, we say the size of the value is N bits. The sizes of all values of a type are always the same, so value sizes are often called as type sizes.

The primitive data types prefixed with "u" are unsigned versions with the same bit sizes. Effectively, this means they cannot store negative numbers, but on the other hand they can store positive numbers twice as large as their signed counterparts. For example, int8 hase range (-128, 127) while uint8 has range (0, 255).

Note:
- `byte` is a built-in alias of `uint8`. We can view `byte` and `uint8` as the same type.
- `rune` is a built-in alias of `int32`. We can view `rune` and `int32` as the same type.

#### 2.2.3 Floating Point Types


#### 2.2.4 Strings and Runes


### 2.7 Type Conversion


### 2.8 var Versus :=


### 2.9 Const


### 2.10 Types and Untyped Constants


## References
https://go101.org/article/basic-types-and-value-literals.html