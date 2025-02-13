```go
func main() {
    m := make(map[string]int)
    m["a"] = 1
    value, ok := m["b"]
    if ok {
        fmt.Println(value)
    } else {
        fmt.Println("Key not found") // Handle the missing key gracefully
    }
}
```