package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negativ number: %d", int(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 1 {
		return x, ErrNegativeSqrt(x)
	} else {
		var z float64 = 1
		var old float64
		for {
			//fmt.Println(z)
			old = z
			z -= (z*z - x) / (2 * z)
			if !isChanging(old, z) {
				return z, nil
			}
		}
	}
}

func isChanging(old float64, value float64) bool {
	var dif float64 = getDif(old, value)
	if old == value {
		return false
	} else if dif < 0.0000001 {
		return false
	} else {
		return true
	}
}

func getDif(old float64, value float64) float64 {
	if old > value {
		return old - value
	} else {
		return value - old
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
