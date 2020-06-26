package main

import "fmt"

import "unsafe"

type hello struct {
	a int8
	b int16
	c int8
}

func main() {
	h := hello{
		a: 1,
		b: 2,
		c: 3,
	}

	fmt.Printf("size: %d \n ", unsafe.Sizeof(h))	//6
}
