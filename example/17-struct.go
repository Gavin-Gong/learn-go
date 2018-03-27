package example

import . "fmt"

// 结构体就是字段集合

// define
type me struct {
	name string
	age  int
}

func init() {
	g := me{"Zen", 24} //简写, 对字段顺序有严格要求
	kf := me{name: "GG", age: 24}
	Println(g, kf)

	// read val
	Println(kf.name)

	// TODO: 遍历?
}
