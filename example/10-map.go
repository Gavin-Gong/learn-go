package example

import (
	. "fmt"
)

func init() {
	// create map, key 只能是非引用类型, value可以是任意类型
	dict := make(map[bool]string) // 声明
	hash := map[string]int{"foo": 1, "bar": 2}
	Println(hash)

	// assign
	dict[true] = "true"
	dict[false] = "false"

	// 获取长度
	Println(len(dict))

	// delet pair of key-value
	delete(dict, true)

	// is Exist
	a, has := dict[true]
	Println(a, has)

	Println(dict)
}
