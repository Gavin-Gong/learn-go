package CircularBuffer

const arraySize = 10 // CB 是确定元素数量的

type CircularBuffer struct {
	data [arraySize]interface{} // 数据
	pointer int // 当前位置
}

// 依次回环插入一个值
func (b *CircularBuffer) InsertValue(i interface{}) {
	if b.pointer == len(b.data) {
		b.pointer = 0
	}
	b.data[b.pointer] = i
	b.pointer++ // 移动光标 以便更新下次的值
}

// 获取所有的数据
func (b *CircularBuffer) GetValues() [arraySize]interface{} {
	return b.data
}

// 制定起点位置， 获取全部数据
// [1,2,3,4,5,6] 指定 3 返回 []
func (b *CircularBuffer) GetValuesFromPos(pos int)([arraySize]interface{}, bool) {
	var res [arraySize]interface{} // 初始化返回结果

	// 数组越界返回错误
	if pos >= len(b.data) {
		return res, false
	}
	// 循环获取数据
	for u := 0; u < len(res); u++ {
		// 抵达最后一位， 重置pos为0
		if pos == len(b.data) {
			pos = 0
		}
		res[u] = b.data[pos]
		pos += 1
	}
	return res, true 
}