package hello

import "fmt"

type VertexID int

func (id VertexID) String() string {
	return fmt.Sprintf("VertexID(%d)", id)
}
