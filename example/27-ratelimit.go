package example

func init() {
	reqs := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		reqs <- i
	}
	close(reqs)
	// lmiters := time.Tick(200 * time.Millisecond)

	// for req
}
