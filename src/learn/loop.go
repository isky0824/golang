package main

import (
	"fmt"
)

func main() {

	//for循环
	for i := 0; i < 10; i++ {
		if i < 3 {
			fmt.Println("Hello world!-----")
			continue
			fmt.Println("-----------------")
		} else if i < 5 {
			fmt.Println("Hello world!+++++")
			break
		} else {
			fmt.Println("Hello world!xxxxx")
		}
	}
}
