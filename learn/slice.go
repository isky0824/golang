/*
 * hello.go
 */

 package main

 import (
	 "fmt"
 )
 
 func main() {
	 fmt.Println("Hello world!")
 
	 slc := []int{1, 2, 3, 4, 5}
	 for k, v range slc {
		 fmt.Println(k,v)
	 }
 }
 