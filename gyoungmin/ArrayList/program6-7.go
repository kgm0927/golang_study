package Arraylist

import (
	"fmt"
	"os"
)

type Element int
type ListNode struct {
	Data Element
	Link *ListNode
}

type Ordinary_ArrayList interface {
	Error(string)
	Insert_first(*ListNode, int)
	Insert(*ListNode, Element) *ListNode
	Delete_first() *ListNode
	Delete(*ListNode) *ListNode
	Print_list()
}

// 리시버로 할 수 있는 방법을 찾지 못하겠음
func (l *ListNode) Error(stderr string) {
	fmt.Println(stderr)
	os.Exit(1)
}

func (l *ListNode) Init_stack() {
	l = &ListNode{Data: 0, Link: nil}
}

func (l *ListNode) Insert_first(head *ListNode, value int) {
	p := new(ListNode)
	p.Data = Element(value)
	p.Link = l
	head = p
	l = head
}

func (l *ListNode) Insert(pre *ListNode, value Element) *ListNode {
	p := new(ListNode)
	p.Data = value
	p.Link = pre.Link
	pre.Link = p
	return l
}

func (l *ListNode) Delete_first() *ListNode {

	if l == nil {
		return nil
	}
	removed := l
	l = removed.Link
	return l

}

func (l *ListNode) Delete(pre *ListNode) *ListNode {

	removed := pre.Link
	pre.Link = removed.Link
	return l

}

func (l *ListNode) Print_list() {
	for p := l; p != nil; p = p.Link {
		fmt.Printf("%d->", p.Data)
	}
	fmt.Printf("NULL \n")
}
