package chap04func

import (
	"fmt"
	"strconv"
	"strings"
)

type BinOp func(int, int) int
type BinSub func(int, int) int
type StrSet map[string]struct{}
type PrecMap map[string]StrSet

func OpThreeAndFour(f BinOp) {
	fmt.Println(f(3, 4))
}

func BinOpToBinSub(f BinOp) BinSub {
	var count int

	return func(i1, i2 int) int {

		fmt.Println(f(i1, i2))
		count++
		return count
	}

}

func Eval(opMap map[string]BinOp, prec PrecMap, expr string) int {

	var ops []string
	var nums []int

	pop := func() int {

		last := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		return last
	}

	reduce := func(higher string) {
		for len(ops) > 0 {
			op := ops[len(ops)-1]

			if strings.Index(higher, op) < 0 {
				return
			}

			ops = ops[:len(ops)-1]

			if op == "(" {
				return
			}

			b, a := pop(), pop()

			if f := opMap[op]; f != nil {
				nums = append(nums, f(a, b))
			}

		}
	}
	for _, token := range strings.Split(expr, " ") {

		switch token {
		case "(":
			ops = append(ops, token)

		case "+", "-":
			reduce("+-*/")
			ops = append(ops, token)

		case "*", "/":
			reduce("*/")
			ops = append(ops, token)

		case ")":
			reduce("+-*/(")

		default:
			num, _ := strconv.Atoi(token)
			nums = append(nums, num)

		}

	}

	reduce("+-*/")
	return nums[0]
}

func NewStrSet(strs ...string) StrSet {
	m := StrSet{}
	for _, str := range strs {
		m[str] = struct{}{}
	}
	return m
}

func NewEvaluator(opMap map[string]BinOp, prec PrecMap) func(expr string) int {

	return func(expr string) int {
		return Eval(opMap, prec, expr)
	}

}
