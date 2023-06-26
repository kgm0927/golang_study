package hangul

import (
	"fmt"
)

func ExampleCount() {
	codeCount := map[rune]int{}
	count("가나다나", codeCount)

	for _, key := range []rune{'가', '나', '다'} {

		fmt.Println(string(key), codeCount[key])

	}
	// Output:
	// 가 1
	// 나 2
	// 다 1
}
