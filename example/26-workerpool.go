package example

import . "fmt"
import "time"

func worker(id int, jobs <-chan int, res chan<- int) {
	for j := range jobs {
		Println("worker", id, "start job", j)
		time.Sleep(time.Second)
		Println("worker", id, "finish job", j)
		res <- j * 2
	}

}

func init() {
	jobs := make(chan int, 100)
	res := make(chan int, 100)

	for i := 0; i <= 10; i++ {
		go worker(i, jobs, res)
	}
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 1; a <= 5; a++ {
		<-res
	}
}
