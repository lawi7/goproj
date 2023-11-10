package main

import "fmt"
import "gmath/gmath"

func main() {
	xs := []float64{1,2,3,4}
	avg := gmath.Average(xs)
	fmt.Println(avg)
}