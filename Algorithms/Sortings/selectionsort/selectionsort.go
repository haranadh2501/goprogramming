package main

import "fmt"

func selectionsort(arr []int) {

	//small := arr[0]

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j <= len(arr)-1; j++ {
			if arr[j] < arr[i] {
				//small = arr[j]
				arr[i], arr[j] = arr[j], arr[i]
			}

		}
		fmt.Println(arr)
	}

}

func main() {

	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(arr)
	fmt.Println("----------Trace Begins------------")
	selectionsort(arr)
	fmt.Println("----------Teace Ended-------------")
	fmt.Println(arr)

}
