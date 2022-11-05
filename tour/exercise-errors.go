package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)	
	}
	z := 1.0
	for i := 0; i < 10; i++ {
		delta := (z*z - x) / (2*z)
		if math.Abs(delta) < 0.000000000001 {
			return z, nil
		}
		z -= delta
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
