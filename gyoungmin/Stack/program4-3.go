package Stack

import (
	"fmt"
	"os"
)

type Element3 int

type StackType struct {
	Data [MAX_STACK_SIZE]Element3
	Top  int
}

func (S *StackType) Init_stack() {
	S.Top = -1
}

func (s *StackType) Is_empty() bool {
	anwr := (s.Top == -1)

	return anwr

}

func (s *StackType) Is_full() bool {
	return (s.Top == (MAX_STACK_SIZE - 1))
}

func (s *StackType) Push(item Element3) {
	if s.Is_full() {
		fmt.Printf("스택 포화 에러")
		return
	} else {
		s.Top += 1 //전위 연산
		s.Data[s.Top] = item
	}
}

func (s *StackType) Pop() Element3 {
	if s.Is_empty() {
		fmt.Println(" 스택 공백 에러")
		os.Exit(1)
	}
	result := s.Top
	s.Top -= 1
	return s.Data[result]

}

func (s *StackType) Peek() Element3 {
	if s.Is_empty() {
		fmt.Println("스택 공백 에러")
		os.Exit(1)
	}

	return s.Data[s.Top]

}
