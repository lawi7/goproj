package main

import "fmt"
import "goproj/src/gobk/c8/math"

func main() {
	xs := []float64{1,2,3,4}
	avg := math.Average(xs)
	fmt.Println(avg)
}