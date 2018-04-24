package example

import . "fmt"
import "time"

func tc() {
	ticker := time.NewTicker(time.Second)
	go func() {
		for t := range ticker.C {
			Println("ticker", t)
		}
	}()
	time.Sleep(5 * time.Second)
	ticker.Stop() // 停止ticker
	Println("stop ticker")
}
