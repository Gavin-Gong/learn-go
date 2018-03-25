package example

import(
	."fmt"
)

func init() {
	var arr [5]int // 未赋值之前都是具有空值的 数字类型就是0
	Println("print arr len %t", len(arr)) // 长度

	// 赋值数组
	assignArr := [...]int{1,2,3,4,5}
	for i := 0; i < len(assignArr); i++ {
		Println(assignArr[i])
	}
}