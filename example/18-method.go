package example

import . "fmt"

type Me struct {
	name string
	age  int
}

// 定义结构体方法, 类似class的方法
// 结构体接受者
func (m Me) ListenMusic() {
	m.name = "music" // 传递的是结构体副本, 修改只在函数内有效
	Println(m.name, "listen music")
}

// 指针结构体接受者
func (m *Me) Code() {
	m.name = "GG" // 传递的是指针, 修改是可以作用到原结构体上去的
	Println(m.name, "code by go")
}

// go 会自动处理指针接受者和结构体接受者之间调用

func mt() {
	me := Me{"zen", 24}
	me.ListenMusic()
	Println(me)
	me.Code()
	Println(me)
}
