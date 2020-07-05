package main

import "fmt"

type A struct {
	name string
	b
}

type B struct {
	name string
}

func main() {
	b := B{
		name: "我是B",
	}
	a := A{
		name: "我是A",
		b:    b,
	}
	fmt.Println("a", a)
}
