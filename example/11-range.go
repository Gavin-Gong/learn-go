package example

import(
	."fmt"
)

func init() {
	// 遍历array/slice
	arr := [...]int{ 1, 2, 3, 4, 5}
	for _, v := range arr {
		Println("print val", v)
	}

	// 遍历 map
	dict := map[string]string{ "key": "value", "k": "v" }
	for k, v := range dict {
		Println(k, v)
	}
	// 只遍历map key
	hashTable := make(map[string]string)
	hashTable["1"] = "1"
	hashTable["2"] = "2"
	for k := range hashTable {
		Println(k)
	}

	// 遍历 string, i -> index, c -> unicode 码点
	for i, c := range "golang" {
		Println(i, c)
	}


}