package example

import (
	. "fmt"
	"time"
)

func tm() {
	timer := time.NewTimer(2 * time.Second) // define a timer
	<-timer.C                               // 通過通道傳遞信息
	Println("start")

	timer2 := time.NewTimer(time.Second)
	go func() {
		// TDOO:不執行是什麽情況
		<-timer2.C
		Println("expire")
	}()
	stop := timer2.Stop()
	if stop {
		Println("stop")
	}
}
