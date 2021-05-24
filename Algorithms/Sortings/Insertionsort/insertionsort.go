package main

import "fmt"

func insertionsort(arr []int) {

	var i, j, v int
	for i = 1; i <= len(arr)-1; i++ {

		v = arr[i]
		j = i
		//fmt.Println(i, j, v)
		for (j >= 1) && (arr[j-1] > v) {
			arr[j] = arr[j-1]
			j--
			// be careful j>=1 should be evaluated first.
			// if arr[j-1] > v gets eval`ed first then out of bounds exception
			// GO evaluates expression from R to L
		}
		arr[j] = v
		fmt.Println(arr)

	}

}

func main() {

	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(arr)
	fmt.Println("----------Trace Begins------------")
	insertionsort(arr)
	fmt.Println("----------Teace Ended-------------")
	fmt.Println(arr)

}
