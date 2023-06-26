package chap04func

import "fmt"

type VertexID int
type EdgeID int

func NewVertexIDGenerator() func() VertexID {
	var next int

	return func() VertexID {
		next++
		return VertexID(next)

	}

}

func NewEdgeIDGenerator() func() EdgeID {

	var next int
	return func() EdgeID {
		next++
		return EdgeID(next)
	}

}

func VertexID_print() {
	i := VertexID(100)
	fmt.Println(i)

}

func (id VertexID) String() string {
	return fmt.Sprintf("Vertex(%d)", id)
}

func VertexID_String() {
	i := VertexID(100)
	fmt.Println(i)

}
