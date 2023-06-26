package concurrency

import "fmt"

func concurrency_debug() {

	req, resp := make(chan struct{}), make(chan int64)
	// resp
	cnt := int64(10)
	go func(cnt int64) {

		defer close(resp)

		for _ = range req {
			cnt--
			resp <- cnt
		}
	}(cnt)

	for i := 0; i < 10; i++ {

		go func() {
			// do something
			req <- struct{}{}
		}()

	}

	for cnt = <-resp; cnt > 0; cnt = <-resp {

	}
	close(req)
	fmt.Println(cnt)

}
