package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, int) {
	//note different declarations method for Go
	//alternative to the declaration below would be
	//var z float64 = (x/2)
	z := float64(x / 2)
	i := 0
	delta := float64(0.000000000001)
	
	//for loop is similar to the one in C, but can be turned
	// into a while loop. Just delete the  2 ';', and 
	//delete initializations and incrementations and you have a
	//free 'while' loop
	for v := math.Abs(z*z - x); v > delta; {
		z -= ((z * z) - x) / (2 * z)
		i++
		v = math.Abs(z*z - x)
	}
	return z, i
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
