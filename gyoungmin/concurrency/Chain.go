package concurrency

import "context"

type IntPipe func(ctx context.Context, _ <-chan int) <-chan int

/*func Chain(ps ...IntPipe) IntPipe {
	return func(in <-chan int) <-chan int {
		c := in
		for _, p := range ps {
			c = p(c)
		}
		return c
	}
}
*/
