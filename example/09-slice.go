package example

func init() {
	// 创建
	slice := make([]string, 3)

	// 赋值
	slice[0] = "no1"

	// 扩展
	slice = append(slice, "ss")
	slice = append(slice, "ss", "nedd more")

	// copy slice
	copy_slice := make([]string, len(slice))
	copy(copy_slice, slice)
}
