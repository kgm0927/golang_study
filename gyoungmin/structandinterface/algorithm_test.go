package structandinterface

import (
	"container/heap"
	"fmt"
)

func (c *CaseInsensitive) Push(x interface{}) {
	*c = append(*c, x.(string))
}

func (c *CaseInsensitive) Pop() interface{} {
	len := c.Len()
	last := (*c)[len-1]
	*c = (*c)[:len-1]
	return last
}

func ExampleCaseInsensitive_heap() {
	apple := CaseInsensitive([]string{"iPhone", "iPad", "MacBook", "AppStore"})
	heap.Init(&apple)

	for apple.Len() > 0 {
		popped := heap.Pop(&apple)
		s := popped.(string)
		fmt.Println(s)
	}
	// Output:
	// .
}
