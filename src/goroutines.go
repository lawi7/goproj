package main

import (
	"fmt"
	"time"
)

func myPrint(start, finish int) {
	for i := start; i <= finish; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println(100 * time.Microsecond)
}

func main() {
	for i := 0; i < 5; i++ {
		go myPrint(i, 5)
	}
	time.Sleep(time.Second)
}