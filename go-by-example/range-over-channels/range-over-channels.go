package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "msg1"
	queue <- "msg2"
	//当没有监听chan且生产者一直往chan写数据时，超过chan大小容量则会发生deadlock
	queue <- "msg3"
	close(queue)

	for data := range queue {
		fmt.Println(data)
	}

}
