package main

import (
	"fmt"
	"time"
)

//-------------------goroutine示例:
func log(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

//通道Channel用于goroutine之间传递数据
//-------------------通道Channel示例:不带缓冲区的通道
func channelWithoutCache0(c chan int) {
	c <- 100
}

func channelWithoutCache1(c chan int) {
	v := <-c
	fmt.Println(v)
	fmt.Println("x")
}

//-------------------通道Channel示例:带缓冲区的通道
func channelWithCache(c chan int, len int) {
	for i := 0; i < len; i++ {
		c <- (i * 2)
	}
	close(c) //发送完数据之后，关闭通道
}

func main() {

	//sample 0
	go log("Hello")
	log("World")

	//sample 1
	/*c := make(chan int) //创建一个命名为c的通道，通道传递的数据是一个int
	go channelWithoutCache0(c)
	go channelWithoutCache1(c)
	v := <-c
	fmt.Println(v)*/

	//sample 2
	/*d := make(chan int, 5)    	//创建一个带缓冲区的通道，缓冲区大小是5个int
	go channelWithCache(d, cap(d)) 	//cap返回缓冲区大小

	//遍历从通道中接收到的数据
	for i := range d {
		fmt.Println(i)
	}*/
}
