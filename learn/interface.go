/*
	// 定义接口
	type interface_name interface {
		method_name1 [return_type]
		method_name2 [return_type]
		method_name3 [return_type]
		...
		method_namen [return_type]
	}

	// 定义结构体
	type struct_name struct {
		// variables
	}

	// 实现接口方法
	func (struct_name_variable struct_name) method_name1() [return_type] {
		// 方法实现
	}
	...
	func (struct_name_variable struct_name) method_namen() [return_type] {
		// 方法实现
	}
*/

package main

import (
	"fmt"
)

// 定义phone接口
type Phone interface {
	call()
}

// 定义NokiaPhone类去实现phone接口
type NokiaPhone struct {
}

// 实现接口中的call方法
func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

// 定义IPhone类去实现phone接口
type IPhone struct {
}

// 实现接口中的call方法(将call方法关联到IPhone类型)
func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main() {
	var phone Phone
	i_xx := 10
	fmt.Println(i_xx)
	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

}
