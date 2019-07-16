# Arrays
Syntax for declaring an array in Go is as follows:
```
var arrayname [size] type
```

You can then assign a value to any element of the array using the index as follows
```
arrayname[index] = value
```

you can initialize an array during the declaration using below syntax
```
arrayname := [size] type {value_0,value_1,…,value_size-1} 
```

You can make the compiler pick the length of the array b using ```[...]``` in place of the [size]
```
arrayname = [...] type {value_0,value_1,…,value_size-1} 
```

You can get the length of the array by using the length() function
```
len(arrayname)
```

#Slicing
Similar to Python, Go provices various ways to work with slicing arrays. A slice is a portion or a segment of an array. The basic syntax for slicing an array is as follows.

```
var slice_name [] type = array_name[start:end]
```
This will create a slice named slice_name from an array named array_name with the elements at the index start to end-1.

You can use following functions to manipulate slices
```
len(slice_name) - returns the length of the slice

append(slice_name, value_1, value_2) - It is used to append value_1 and value_2 to an existing slice.

append(slice_nale1,slice_name2…) – appends slice_name2 to slice_name1
```

