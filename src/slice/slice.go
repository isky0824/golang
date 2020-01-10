/*
 * hello.go
 */

package main

import (
	"fmt"
)

type Cat struct {
	Color string
	Name  string
}

type BlackCat struct {
	Cat // 嵌入Cat, 类似于派生
}

// “构造基类”
func NewCat(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

// “构造子类”
func NewBlackCat(color string) *BlackCat {
	cat := &BlackCat{}
	cat.Color = color
	return cat
}

func test(x, y int) (z int) {
	z = x + y
	return
}

func fire(int) {
	fmt.Println("fire")
}

func main() {
	fmt.Println("Hello world!")
	//var err = errors.New("this is an error")
	//fmt.Println(type(err))

	ketty := NewBlackCat("Black")
	fmt.Printf("%v", ketty)
}
