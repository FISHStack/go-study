package main

import "fmt"

import "unsafe"

type persion1 struct {
	a bool
	c int32
	b int8
	d int64
	e byte
	// f string
}

type persion2 struct {
	e byte
	b int8
	a bool
	c int32
	d int64
	// f string
}

func main() {

	fmt.Printf("bool align: %d\n", unsafe.Alignof(bool(true)))         //1
	fmt.Printf("int32 align: %d\n", unsafe.Alignof(int32(0)))          //4
	fmt.Printf("int8 align: %d\n", unsafe.Alignof(int8(0)))            //1
	fmt.Printf("int64 align: %d\n", unsafe.Alignof(int64(0)))          //8
	fmt.Printf("byte align: %d\n", unsafe.Alignof(byte(0)))            //1
	fmt.Printf("string align: %d\n", unsafe.Alignof("EDDYCJY"))        //8
	fmt.Printf("map align: %d\n", unsafe.Alignof(map[string]string{})) //8

	fmt.Println("****************************")

	p1 := persion1{}
	p2 := persion2{}

	fmt.Printf("part1 size: %d , align: %d \n", unsafe.Sizeof(p1), unsafe.Alignof(p1)) //32 8

	fmt.Printf("part2 size: %d , align: %d \n", unsafe.Sizeof(p2), unsafe.Alignof(p2))                    //16 8
	fmt.Printf("string size: %d , align: %d \n", unsafe.Sizeof(string("a")), unsafe.Alignof(string("a"))) //16 8
	fmt.Printf("int8 size: %d , align: %d \n", unsafe.Sizeof(int8(127)), unsafe.Alignof(int8(127)))       //1 1
}
