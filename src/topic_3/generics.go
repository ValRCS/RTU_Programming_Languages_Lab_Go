package main

import "fmt"

// Ordered constraint: types that support < and >
// This avoids external libraries and keeps the example simple.
type Ordered interface {
    ~int | ~int32 | ~int64 | ~float64 | ~string
}

// Max returns the larger of two ordered values.
func Max[T Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}

func main() {
    fmt.Println("Max(3, 7) =", Max(3, 7))
    fmt.Println(`Max("go", "rtu") =`, Max("go", "rtu"))
}
