package main

import "fmt"

func main() {
	arr := []int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90}
	var ele int
	fmt.Scanf("%d", &ele)
	fmt.Println("")
	var index int
	var found bool
	for index = 0; index < len(arr); index++ {
		if arr[index] == ele {
			fmt.Printf("the element %v is found at the index %v", ele, index)
			found = true
			break
		}
	}
	if found == false {
		fmt.Println("The entered element is not in the array.")
	}
}
