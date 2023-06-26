package main

import "fmt"

func MergeSort(array []int, p, k int) {

	MergeSort(array, p, k/2)
	MergeSort(array, (k/2)+1, k)

	if k-p != 1 {
		if p == array[p] {
			fmt.Printf("%d 찾았습니다!", p)
		}

	}

	if k-p != 1 {
		if k == array[k] {
			fmt.Printf("%d 찾았습니다.", p)
		}
	}
	fmt.Println("맞는 것을 찾을 수 없습니다.")

}

func main() {

	array := []int{2, 4, 6, 3, 4, 7, 8, 9}

	MergeSort(array, 0, 7)

}
