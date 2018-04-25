package example

import (
	"flag"
	"fmt"
	"os"
)

// 基础命令行使用
func basicCL() {
	argWithProg := os.Args // 第一个参数是程序名称
	argWithoutProg := os.Args[1:]
	arg := os.Args[2] // 取到某个参数

	fmt.Println(argWithProg)
	fmt.Println(argWithoutProg)
	fmt.Println(arg)
}

// 命令行标志用法
// 基本的标志声明只支持字符串 整数 布尔
//
func flagCL() {
	wordPt := flag.String("word", "foo", "string desc")
	numPt := flag.Int("num", 64, "num desc")
	boolPt := flag.Bool("bool", true, "bool desc")

	// 另一种声明方法
	var svar string
	flag.StringVar(&svar, "other", "str", "other desc")

	flag.Parse() // 必须在定义参数之后
	fmt.Println("word", *wordPt)
	fmt.Println("num", *numPt)
	fmt.Println("bool", *boolPt)
	fmt.Println("other", svar)
	fmt.Println("tail", flag.Args())

}
func init() {
	// basicCL()
	flagCL()
}
