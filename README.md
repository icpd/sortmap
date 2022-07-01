# sortmap

# Usage
```go
m := sortmap.New[uint]()
m.Set(2, "a")
m.Set(1, "b")

m.Keys() // 2,1

m.Get(1) // b
```
