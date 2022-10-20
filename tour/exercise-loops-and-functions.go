package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		delta := (z*z - x) / (2*z)
		if math.Abs(delta) < 0.000000000001 {
			break
		}
		z -= delta
		fmt.Println(i, z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
