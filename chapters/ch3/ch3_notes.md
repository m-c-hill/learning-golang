# Chapter 3: Composite Types

## 3.0 Arrays

Arrays are rarely used directly in Go.

All elements in an array must of the type that has been specified. Declaration:

```
var x [3]int  // All positions initialised to zero value

var x [3]int{10, 20, 30}

var x [12]int{1, 5: 4, 6, 10: 100, 15}  // Creates array [1, 0, 0, 0, 0, 4, 6, 0, 0, 0, 100, 15]
```

In the final example, `5:` declares that all following elements up to index 5 are set to the zero value.

Go only has one-dimensional arrays, but multidimenstional arrays can be simutaled:

`var x [2][3]int`

Here, x is an array of length 2 whose type is an array of ints of length 3.

Built in function `len` to return array length.

Limitation: size of an array is part of the 'type'. You cannot use a type conversion to convert arrays of different sizes into identical types, so arrays functions can only work on one size and you can't assign arrays of different sizes to the same variable.

A variable cannot be used to specify the size of an array as types are reoslved at compile time, not runtime.

## 3.1 Slices
Slices are used more often than arrays in situations where you want a data structure to hold a sequence of values.

Unlike arrays, the length of the slice is *not* part of the slice's type (ie. we can write slice functions that can process any size of slice as well as grow slices when needed).

A slice does not store any data, it just describes a section of an underlying array.

### 3.1.1 Slice Declaration

Declaration:
`var x = []int{10, 20, 30}  // [...] makes an array, [] makes a slice`

To grow slices, use `append`:

```
var x []int{1, 0, 4}
x = append(x, 10) // append single element
x = append(x, 1, 2, 3) // append multiple elements
```

Each element in a slice is assigned to consecutive memory locations (quick to read or write). Each slice has a capacity assigned to it - this is the number of consecutive memory locations reserved for it to use (which can be larger than the length). Use `cap` function to return the slice's current capacity.

Go slices will double in capacity when required (if cap < 1024). Once it's over this limit, it will grow by at least 25%.

The `make` function allows us to declare a slice while also specifying its type, length and capacity. It's advantagous to create a slice with the correct initial capacity if you know ahead of time that it will grow to a certain length. You want to minimise the number of times a slice needs to grow.

```
x := make([]int, 5) // slice with length 5, capacity 5 (by default)
x := make([]int, 5, 10) // slice with length 5, capacity 10 (specified)
```

## 3.1.2 Slicing Slices

Slice expressions will create a slice from a slice using `[<start>:<stop>]`, similar to Python.
```
x := []int{1, 2, 3, 4}
y := x[:2]  // prints [1 2]
z := x[1:]  // prints [2 3 4]
d := x[1:3] // prints[2 3]
```

Taking a slice from a slice is _not_ making a copy of the data by instead creating a second variable that will share memory with the original slice. Changes to a slice of a slice will affect the original slice.

To avoid complicated slice sitations, never use `append` with a sub-slice, or make sure that `append` doesn't cause an overwrite by using a _full slice expression_.
```
y := x[:2:2]
z := x[2:4:4]
```

Above, both y and z have a capacity of 2, where `capacity = max - low`:

`a[low : high : max]`

### 3.1.3 Copying a Slice
Use `copy` to create a slice that's independent of the original: `copy(<destination_slice>, <source_slice>)`

```
x := int{1, 2, 3, 4}
y := make([]int, 4)
num := copy(y, x)
fmt.Println(y, num) // [1 2 3 4] 4
```

## 3.2 The Go Runtime

Go runtime provides services such as memory allocation, garbage collection, concurrency support and networking. Unlike other language, the Go runtime is compiled into every Go binary (avoid use of vm), makes distribution of Go programs easier.


## 3.3 Strings, Runes and Bytes

### 3.3.1 Strings

Go uses a sequence of bytes to represent a string (often assumed to be UTF-8 encoded).

Strings can also be indexed, much like slices:
```
var s string = "Hello there"
var b byte = s[6]  // b set to 't'
var s2 string = s[:5] // "Hello"
```

Warning: some UTF-8 characters can be represented by multiple bytes (such as emojis) which can lead to problems when slicing.

### 3.3.2 String conversions

A single run or byte can be converted to a string:

```
var a rune = 'x'
var s string = string(a)
var b byte = 'y'
var s2 string = string(b)
```

A string can be converted back and forth to a slice of bytes or a slice of runes:

```
var s string = "Hello"
var bs []byte = []bytes(s)
var rs []rune = []rune(s)
```

## 3.4 Maps



## References
- https://go101.org/article/basic-types-and-value-literals.html
- https://stackoverflow.com/questions/27938177/golang-slice-slicing-a-slice-with-sliceabc