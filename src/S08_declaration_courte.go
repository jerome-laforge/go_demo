package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// Ne peut pas utiliser v ici.
	return lim
}

func main() {
	fmt.Println(pow(2, 2, 10))
	fmt.Println(pow(2, 4, 10))
}
