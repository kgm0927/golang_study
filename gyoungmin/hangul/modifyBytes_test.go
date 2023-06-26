package hangul

import "fmt"

func Example_modifyBytes() {
	b := []byte("가나다")
	fmt.Println(b)
	b[2]++
	fmt.Println(b)
	//Output:
	//.
}
