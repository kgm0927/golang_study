package hangul

import "fmt"

func Example_slicing() {
	nums := make([]int, 10)
	nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 13, 14, 15}
	fmt.Println(nums)
	fmt.Println(nums[1:3])
	fmt.Println((nums[2:]))
	fmt.Println(nums[:3])
	//Output:
	//.

}
