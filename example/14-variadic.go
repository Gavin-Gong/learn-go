package example

import (
	. "fmt"
)

func init() {
	arr := []int{1, 2, 3, 4, 5} // 无法使用array
	Println(sum(arr...))        // 字符串解构 我的天
}

// 定義函數時候的可變參數
func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}
