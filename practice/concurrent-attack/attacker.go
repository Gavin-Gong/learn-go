package attacker

import (
	"fmt"
	"net/http"
)

func Start() {
	// channel := make(chan string, 1000)
	for i := 0; i < 1000; i++ {
		go func() {
			resp, err := http.Get("http://8l.ee/")
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(i, resp.StatusCode)
		}()
	}
}
