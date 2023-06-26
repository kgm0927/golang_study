package Stack

import (
	"fmt"
	"os"
)

type Element5 rune

type StackType3 struct {
	Data [MAX_STACK_SIZE]Element5
	Top  int
}

func (S *StackType3) Init_stack() {
	S.Top = -1
}

func (s *StackType3) Is_empty() bool {
	anwr := (s.Top == -1)

	return anwr

}

func (s *StackType3) Is_full() bool {
	return (s.Top == (MAX_STACK_SIZE - 1))
}

func (s *StackType3) Push(item Element5) {
	if s.Is_full() {
		fmt.Printf("스택 포화 에러")
		return
	} else {
		s.Top += 1 //전위 연산
		s.Data[s.Top] = item
	}
}

func (s *StackType3) Pop() Element5 {
	if s.Is_empty() {
		fmt.Println(" 스택 공백 에러")
		os.Exit(1)
	}
	result := s.Top
	s.Top -= 1
	return s.Data[result]

}

func (s *StackType3) Peek() Element5 {
	if s.Is_empty() {
		fmt.Println("스택 공백 에러")
		os.Exit(1)
	}

	return s.Data[s.Top]

}

func (e *StackType3) Prec(op Element5) int {

	switch {

	case op == ')' || op == '(':
		return 0
	case op == '+' || op == '-':
		return 1
	case op == '*' || op == '/':
		return 2

	}

	return -1
}

func Infix_to_postfix(exp string) {

	i := 0

	var ch, top_op Element5
	len := len(exp)

	var s StackType3
	s.Init_stack() // 스택 초기화

	for i = 0; i < len; i++ {
		ch = Element5(exp[i])

		switch {

		case ch == '+' || ch == '-' || ch == '*' || ch == '/': // 연산자
			for (!s.Is_empty()) && (s.Prec(ch) <= s.Prec(s.Peek())) {
				fmt.Printf("%c", s.Pop())
			}
			s.Push(ch)

		case ch == '(':

			s.Push(ch)

		case ch == ')':
			top_op = s.Pop()

			for top_op != '(' {
				fmt.Printf("%c", top_op)
				top_op = s.Pop()
			}

		default:
			fmt.Printf("%c", ch)

		}

	}

	for !s.Is_empty() {
		fmt.Printf("%c", s.Pop())
	}
}

// ////////////////////////////////////////////////////////////////
func (s *StackType3) Insert(exp string) {

	for i, v := range exp {

		s.Data[i] = Element5(v)

	}
	for i, v := range s.Data {

		r := string(v)
		fmt.Println(i, v, r)
	}

}

/*
switch는 다른 언어와 다르게 break가 내재되어 있으므로 따로 쓸 필요가 없으며, case를 연결해서 쓸 수도 없다.

*/
