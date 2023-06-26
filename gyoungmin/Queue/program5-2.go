package Queue

import (
	"fmt"
	"os"
)

type Element2 int

type QueueType2 struct {
	Front int
	Rear  int
	Data  [MAX_QUEUE_SIZE]Element2
}

type Combine interface {
	Qu
	CircleQu
}

type CircleQu interface {
	Peek() Element2
}

func (q *QueueType2) Error(stderr string) {
	fmt.Println(stderr)
	os.Exit(1)
}

func (q *QueueType2) Init_queue() {
	q.Front = 0
	q.Rear = 0
}

func (q *QueueType2) Is_full() bool {
	result := (((q.Rear + 1) % MAX_QUEUE_SIZE) == q.Front)
	return result
}

func (q *QueueType2) Is_empty() bool {
	return q.Front == q.Rear
}

func (q *QueueType2) Enqueue(item int) {
	if q.Is_full() {
		q.Error("큐가 포화상태입니다.")
	}
	q.Rear = ((q.Rear + 1) % MAX_QUEUE_SIZE)

	q.Data[q.Rear] = Element2(item)
}

func (q *QueueType2) Dequeue() int {
	if q.Is_empty() {
		q.Error("큐가 공백상태입니다.")
	}
	q.Front = ((q.Front + 1) % MAX_QUEUE_SIZE)
	return int(q.Data[q.Front])
}

func (q *QueueType2) Queue_print() {

	fmt.Printf("QUEUE(front=%d rear=%d) = ", q.Front, q.Rear)

	if !q.Is_empty() {

		i := q.Front

		for i != q.Front {

			i = (i + 1) % MAX_QUEUE_SIZE
			fmt.Printf("%d | ", q.Data[i])

			if i == q.Rear {
				break
			}
		}
	}
	fmt.Println()
}

func (q *QueueType2) Peek() Element2 {
	i := Element2(0)
	return i
}

//보류
