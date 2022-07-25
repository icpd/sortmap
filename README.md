# Sortmap
Sortmap is map where the keys keep the order that they're added.

# Usage
```go
m := sortmap.New[int,string]()
m.Set(2, "a")
m.Set(1, "b")

m.Keys() // 2,1

m.Get(1) // b
```
