package example

import (
	. "fmt"
)

func cs() {
	// 匿名函数/ 立即执行函数
	func() {
		call := closures()
		Println(call(), call(), call())
	}()
}

// 闭包
func closures() func() int {
	i := 1
	return func() int {
		i++
		return i
	}
}
