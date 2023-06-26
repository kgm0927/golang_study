package concurrency

import (
	"fmt"
	"time"
)

func ten_gorutine() {

	cnt := int64(10)

	for i := 0; i < 10; i++ {

		go func() {
			// do something
			cnt--
		}()
	}

	for cnt > 0 {
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println(cnt)
}
