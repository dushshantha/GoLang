# Conditions
Conditions in Go is little simplied than most languages and very close to Python

```
if condition{
// statements_1
}else{
// statements_2
}
```

A Chained if-else block is as simple as below.

```
if condition_1{
// statements_1
}else if condition_2{
// statements_2
} else {
    statement_3
}
```

Go also supports more traditional Switch statements. The syntax is something like this:

```
switch expression {
    case value_1:
        statements_1
    case value_2:
        statements_2
    case value_n:
        statements_n
    default:
        statements_default
    }
```