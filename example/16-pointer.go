package example

import . "fmt"

func pt() {
	val := 1

	// & 取指针地址
	p := &val
	// * 取指针值
	Println(p, *p)

}
