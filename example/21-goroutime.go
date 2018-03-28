package example

import . "fmt"

/*
 * goroutine 本质上是一个轻量级的线程, 线程和进程的区别?
 * goroutime是go支持高并发的核心所在
 * goroutine 执行的函数会与非goroutine代码同时执行
 */
func init() {
	console("no goroutine")
	go func(str string) {
		console(str)
	}("goroutine")
	console("no goroutine")
}

func console(s string) {
	for i := 0; i < 8; i++ {
		Println(s, i)
	}
}
