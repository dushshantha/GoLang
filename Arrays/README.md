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
