package concurrency

import "fmt"

func Example_simpleChannel() {

	ch := func() <-chan int { //받기 전용
		c := make(chan int)
		go func() {

			defer close(c)
			c <- 1
			c <- 2
			c <- 3
		}()
		return c
	}()

	for num := range ch {
		fmt.Println(num)
	}

	//Output:
	// 1
	// 2
	// 3

}
