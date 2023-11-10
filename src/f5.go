package main

import (
	"fmt"
	"hash/crc32"
	"crypto/sha1"
)

func main() {
	//create a hasher
	h := crc32.NewIEEE()
	g := sha1.New()

	// write data to it
	h.Write([]byte("testing"))
    g.Write([]byte("testing"))

	// calculate the crc32 hash
	v := h.Sum32()
	w := g.Sum([]byte{})

	fmt.Println(v)
	fmt.Println(w)
}