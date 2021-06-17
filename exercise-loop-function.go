package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 0.5
	for i := 0; i < 10; i++ {
		tmp := z
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
		if math.Abs(tmp-z) < math.Pow(10, -15) {
			break
		}
	}

	return z
}

func main() {
	fmt.Println("the close is :")
	fmt.Println(Sqrt(3) - math.Sqrt(3))
}
