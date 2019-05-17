# Set-in-Go
This repository gives the implementation of set, a kind of data structure, of which the elements are non-repeated in Go.
The set is implemented by a map, of which the size can be infinity, with methods: NewSet(), Add(), Remove(), Contains(), Len(), Clear(), IsEmpty(), List().
The set structure is thread safe.

Notes:
As for method List(), the element returned is []interface{} type. If you want to transform the element of list into type int, type string and other types, write the following code in your program:
```go
switch v := element.(type) {
    case int:
        ...
    case string:
        ...
    ...
}
```
