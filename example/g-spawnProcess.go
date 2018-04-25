package example

import (
	"fmt"
	"os/exec"
)

func init() {
	spawnProcess()
}

func spawnProcess() {
	// 相当于执行 vue -V命令
	dateCmd := exec.Command("vue", "-V")

	// 输出vue命令的结果
	dateOut, err := dateCmd.Output()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("-> Vue")
	fmt.Println(string(dateOut))
}
