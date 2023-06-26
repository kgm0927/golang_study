package Stack

import (
	"fmt"
	"os"
)

const MAX_STACK_SIZE = 100

type Element int

var Stack [MAX_STACK_SIZE]Element

var top int

func Init() {
	top = -1
}

func Is_empty() bool {

	return (top == -1)

}

func Is_full() bool {
	return (top == (MAX_STACK_SIZE - 1))
}

func Push(item Element) {

	if Is_full() {
		fmt.Print("스택 공식 에러 \n")
		return
	} else {
		top++
		Stack[top] = item

	}
}

func Pop() Element {
	var fix int
	if Is_empty() {

		fmt.Print("스택 공식 에러 \n")
		os.Exit(1)
	} else {
		fix = top
		top--
	}
	return Stack[fix]
}

func Peek() Element {
	if Is_empty() {
		fmt.Print("스택 공식 에러\n")
		os.Exit(1)
	}

	return Stack[top]
}
