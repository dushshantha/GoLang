# Variables (and constants) of Go
This simple code demonstrates use of the variables and constants in a Go program

Variable declaration follows few different patters in Go.

There are 3 basic Types available in Go.
1. Numeric Types : various numerics types which include ints, floats and complex values. [Types can be found here](https://golang.org/search?q=Types#Variables)
2. String types : a sequence of bytes that you can perform various operations like concatantion, substring.
3. Boolean types : true or false

You can declare a variable with following syntax:
```
var <variablename> <type>
```

You can declare a variable with an initial value using below syntax:
```
var <variablename> <type> = <value>
```
(Here, the value should be a valid value for the type)

You can also let Go infer the type from an initial value by using below syntax:
```
var <variablename> = <value>
```

You can also declae multiple variables in one line:
```
var <variablename1>, <variablename2> = <value1>, <value2>
```

And finally, You can declare and assign a variable without using var keyword as follows. Here, please note the use of ":=" notation
```
<variablename> := <value>
```

