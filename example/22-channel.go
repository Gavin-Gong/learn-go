package example

import (
	"fmt"
	"time"
)

/*
 * channel 是通信管道 用于goruotine之间传递数据
 */
func init() {
	c := make(chan string) // 新建一个 channel
	go func(c chan string) {
		time.Sleep(1000 * time.Millisecond)
		c <- "from channel" // 发送数据到channel
	}(c)
	msg := <-c // channel接收数据, 阻塞后面代码的执行
	fmt.Println(msg)

	// buffer channel 处理传递多个值得情况
	ch := make(chan string)
	ch <- "1"
	ch <- "2"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
