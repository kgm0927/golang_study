package concurrency

import "context"

// PlusOne returns a channel of num+1 for nums received from in.

func PlusOne(ctx context.Context, in <-chan int) <-chan int {

	out := make(chan int)

	go func() {
		defer close(out)
		for num := range in {

			select {
			case out <- num + 1:

			case <-ctx.Done():
				return

			}
		}
	}()
	return out
}

/*
func ExamplePlusOne() {
	c := make(chan int)
	go func() {
		defer close(c)
		c <- 5
		c <- 3
		c <- 8

	}()
	for num := range PlusOne(PlusOne(c)) {
		fmt.Println(num)
	}

}
*/
