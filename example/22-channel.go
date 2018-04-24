package example

import (
	"fmt"
	"time"
)

/*
 * channel 是通信管道 用于goruotine之间传递数据
 */
func cn() {
	c := make(chan string) // 新建一个 channel
	go func(c chan string) {
		time.Sleep(1000 * time.Millisecond)
		c <- "from channel" // 发送数据到channel
	}(c)
	msg := <-c // channel接收数据, 阻塞后面代码的执行
	fmt.Println(msg)

	// buffer channel 处理传递多个值得情况,
	ch := make(chan string, 2)
	ch <- "1"
	ch <- "2"
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// 我们可以指定 通道是发送数据合适接受数据 以此来提高类型安全性
	// func ping(r <-chan string, g chan<- string){
	// msg :=
	// }
}

// 通道选择器
func slct() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		c1 <- "One"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "Two"
	}()
	// 必须循环一下
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

// 关闭通道
// 关闭一个通道意味着不能再向这个通道发送值了
func closeChan() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("get", j)
			} else {
				fmt.Println("get all")
				done <- true
				return
			}
		}
	}()
	for i := 0; i < 3; i++ {
		fmt.Println("send", i)
		jobs <- i // 发送
	}
	close(jobs)
	<-done
}

// 遍历缓冲通道， 一个非空通道也是可以关闭的， 剩下的值任然可以接受到
func rangeChan() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)
	for ele := range queue {
		fmt.Println(ele)
	}
}

func init() {
	// cn()
	// slct()
	// closeChan()
	// rangeChan()
}
