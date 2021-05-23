package main

import "fmt"

func bubblesort(arr []int) {

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {

	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(arr)
	bubblesort(arr)
	fmt.Println(arr)

}
