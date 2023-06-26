package concurrency

import (
	"fmt"
	"sync"
)

func Once_sync() {
	var Once sync.Once
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			Once.Do(func() {
				fmt.Println("Initialized")
			})
			fmt.Println("Goroutine:", i)
		}(i)
	}
	wg.Wait()
}
