package main

import (
	"fmt"
	"runtime"
	"sync"
)

//当前打印的总票数
var count = 0

//需要打印的总票数
var total = 5

//同步锁
var lock sync.Mutex

//等待
var wg sync.WaitGroup

//打印车票的协程
func test() {
	//稍后协程退出时通知wg
	defer wg.Done()

	//业务逻辑
	for count < total {
		lock.Lock()
		{
			count++
			fmt.Printf("打印第%v张车票\n", count)

			//暂停当前协程，将CPU调度让给其他协程
			runtime.Gosched()
		}
		lock.Unlock()
		//time.Sleep(time.Second)
	}
}

func main() {
	//获取逻辑CPU核心数量，然后设置线程池与CPU核心数的对应关系
	runtime.GOMAXPROCS(runtime.NumCPU())

	//等待2个协程
	wg.Add(2)

	//开启2个协程
	go test()
	go test()

	//主线程阻塞以等待2个协程结束
	wg.Wait()

	//打印结果
	fmt.Printf("count=%v\n", count)
}
