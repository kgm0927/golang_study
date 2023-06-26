package concurrency

import (
	"fmt"
	"sync"
)

func Once() {
	done := make(chan struct{})
	go func() {
		defer close(done)
		fmt.Println("Initialized")
	}()

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {

		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			fmt.Println("Goroutin:", i)
		}(i)
	}
	wg.Wait()
}
