package Queue

import (
	"fmt"
	"os"
)

type Element int

const MAX_QUEUE_SIZE = 5

type QueueType struct {
	Front int
	Rear  int

	Data [MAX_QUEUE_SIZE]Element
}

type Qu interface {
	Error(string)
	Init_queue()
	Queue_print()
	Is_full() bool
	Is_empty() bool
	Enqueue(int)
	Dequeue() int
}

func (q *QueueType) Error(char string) {
	fmt.Println(char)
	os.Exit(1)
}

func (q *QueueType) Init_queue() {
	q.Rear = -1
	q.Front = -1
}

func (q *QueueType) Queue_print() {

	for i := 0; i < MAX_QUEUE_SIZE; i++ {
		if i <= q.Front || i > q.Rear {

			fmt.Print("   | ")

		} else {

			fmt.Printf("%d | ", q.Data[i])

		}
	}
	fmt.Println()

}

func (q *QueueType) Is_full() bool {
	if q.Rear == MAX_QUEUE_SIZE-1 {
		return true
	} else {
		return false
	}
}

func (q *QueueType) Is_empty() bool {

	if q.Front == q.Rear {
		return true
	} else {
		return false
	}
}

func (q *QueueType) Enqueue(item int) {
	if q.Is_full() {
		q.Error("큐가 포화상태입니다.")
		return
	}
	q.Rear += 1
	change := Element(item)
	q.Data[q.Rear] = change
}

func (q *QueueType) Dequeue() int {
	if q.Is_empty() {
		q.Error("큐가 공백상태입니다.")
		return -1
	}
	q.Front += 1
	item := q.Data[q.Front]
	result := int(item)

	return result
}

/*func main() {

Q := &Queue.QueueType{}
var inter Queue.Qu = Q

for i := 0; i < 3; i++ {
	in := (i + 1) * 10
	inter.Enqueue(in)
	inter.Queue_print()
}
for i := 0; i < 3; i++ {
	_ = inter.Dequeue()
	inter.Queue_print()
}*/
