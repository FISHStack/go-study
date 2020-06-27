package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "msg1"
	}()

	select {
	case msg1 := <-c1:
		fmt.Println("recived:", msg1)
	case <-time.After(1 * time.Second):
		fmt.Println("time out 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "msg2"
	}()

	select {
	case msg2 := <-c2:
		fmt.Println("recived:", msg2)
	case <-time.After(3 * time.Second):
		fmt.Println("time out 2")
	}
}
