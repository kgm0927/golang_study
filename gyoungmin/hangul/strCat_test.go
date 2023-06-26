package hangul

import "fmt"

func Eample_strCat() {
	s := "abc"
	ps := &s
	fmt.Println(s)
	fmt.Println(*ps)
	// Output:
	//abcdef
}
