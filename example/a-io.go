package example

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 讀取當前目錄
func readDir() {
	data, _ := ioutil.ReadDir("./")
	// fmt.Println(data, err)
	for _, v := range data {
		fmt.Println(v.Name())
	}
}

// 讀取文件
func readFile() {
	data, err := ioutil.ReadFile("./temp/file.md")
	check(err)
	fmt.Println(string(data))

	// 更多的讀取控制
	f, err := os.Open("./temp/file.md")
	defer f.Close()
	check(err)

	// 從文件開始讀取字節 但是最多兩個字節
	b1 := make([]byte, 2)
	n1, err := f.Read(b1)
	check(err)
	fmt.Println(string(b1), n1)

	// 跳讀
	o2, err := f.Seek(4, 0)
	check(err)
	b2 := make([]byte, 4)
	n2, err := f.Read(b2)
	check(err)
	fmt.Println(string(b2), n2, o2)

	// bufio 實現了緩衝讀取 有助於提升性能
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(6)
	check(err)
	fmt.Println(string(b4), "xx")
}

func init() {
	// readDir()
	// readFile()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
