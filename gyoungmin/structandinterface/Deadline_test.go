package structandinterface

import (
	"time"
)

type Deadline struct {
	time.Time
}

/*type task struct {
	title    string
	Status   status
	Deadline *Deadline
}*/

func NewDeadline(t time.Time) *Deadline {
	return &Deadline{t}
}

// OverDue 함수는 현 시간에서 입력된 시간이 빠르거나 아니면 지나 있는가를 확인하는 것이다.
