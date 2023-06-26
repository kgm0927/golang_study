package chap04func

import "fmt"

func NewIntGenerator() func() VertexID {
	gen := NewIntGenerator()
	return func() VertexID {
		return VertexID(gen())
	}
}

func ExampleNewIntGenerator() {
	gen := NewIntGenerator()
	gen2 := NewIntGenerator()
	fmt.Println(gen(), gen(), gen(), gen(), gen())
	fmt.Println(gen2(), gen2(), gen2(), gen2(), gen2())
	// Output:
	// 1 2 3 4 5
	// 6 7 8 9 10
}
