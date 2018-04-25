package example

import (
	"fmt"
	"os"
	"strings"
)

func init() {
	// envVar()
}

func envVar() {
	// 设置环境变量
	os.Setenv("API", "https://api.github.com")

	// 获取环境变量
	fmt.Println("API", os.Getenv("API"))

	// 循环输出所有的环境变量
	for _, e := range os.Environ() {
		// e 类似 API=http 的字符串
		pair := strings.Split(e, "=")
		fmt.Println(pair)
	}
}
