package concurrency

import "fmt"

func Example_closedChannel() {
	c := make(chan int)

	close(c)

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	// Output:
	// 0
	// 0
	// 0

}
