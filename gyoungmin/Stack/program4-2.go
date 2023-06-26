package Stack

import (
	"fmt"
	"os"
)

type Element_2 struct {
	Student_no int
	Name       string

	Address string
}

//type Stack_2 [MAX_STACK_SIZE]Element_2 이런 방식으로는 Stack에 할당이 되지 않는다. 그래서 그냥 변수 배열 처리하겠다.

type Stack_2 []Element_2

func (e *Stack_2) Push(item Element_2) {
	if Is_full() {
		fmt.Println("스택 포화에러")
		return

	}
	top += 1
	*e = append(*e, item)

}

func (e Stack_2) Pop() Element_2 {

	if Is_empty() {
		fmt.Printf("스택 공백 에러\n")
		os.Exit(1)
	}

	result := e[top]
	top -= 1

	return result

}

func (e Stack_2) Peek() Element_2 {
	if Is_empty() {
		fmt.Printf("스택 공백 에러\n")
		os.Exit(1)
	}

	result := e[top]

	return result

}
