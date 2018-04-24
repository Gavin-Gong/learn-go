package example

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func writeFile1() {
	d1 := []byte("hello golang")
	err := ioutil.WriteFile("./temp/writeme.md", d1, 06644)
	check(err)
}
func writeFile2() {
	f, err := os.Create("./temp/writeFile2")
	check(err)
	defer f.Close() // 關閉

	d := []byte{115, 115, 116, 117}
	n, err := f.Write(d)
	check(err)
	fmt.Println(n)
}

func writeFile3() {
	f, err := os.Create("./temp/writeFile3")
	defer f.Close()
	check(err)
	n, err := f.WriteString("woowowo") // n 長度
	check(err)
	fmt.Println(n)
}

func bufWrite() {
	f, err := os.Create("./temp/bufWrite")
	defer f.Close()
	check(err)
	f.Sync()
	w := bufio.NewWriter(f)
	n, err := w.WriteString("buffer io")
	fmt.Println(n)
	check(err)
	w.Flush()
}

func init() {
	// writeFile3()
	// bufWrite()
}
