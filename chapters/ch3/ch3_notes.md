# Chapter 3: Composite Types

### 3.0 Arrays

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

### 3.1 Slices



## References
https://go101.org/article/basic-types-and-value-literals.html

