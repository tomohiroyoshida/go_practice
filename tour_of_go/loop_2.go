package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
    z := 1.0
    for d := 1.0; d*d > 1e-10; z -= d {
        d = (z*z - x) / (2 * z)
    }
    return z
}

func main() {
    fmt.Println(Sqrt(2))
}