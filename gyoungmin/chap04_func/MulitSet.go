package chap04func

type MultiSet map[string]int
type SetOp func(m MultiSet, val string)

func Insert(m MultiSet, val string) {
	m[val]++
}

func (m MultiSet) Erase(val string) {
	if m[val] <= 1 {
		delete(m, val)
	} else {
		m[val]--
	}
}

func (m MultiSet) Count(val string) int {
	return m[val]
}
