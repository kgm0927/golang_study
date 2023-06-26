package structandinterface

import "testing"

type status int
type ByteSize float64

const (
	UNKNOWN status = iota
	TODO    status = iota
	DONE    status = iota
)

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func AssertIntEqual(t *testing.T, a, b int) {
	if a != b {
		t.Errorf("assertion failed: %d!=%d", a, b)
	}

}
