package main

import "fmt"

func main() {
	arr := []int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90}
	var k int
	fmt.Scanf("%d", &k)
	fmt.Println("")

	for i := 0; i < k; i++ {
		for j := 0; j < len(arr)-2-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		fmt.Printf("%v large element is %v\n", i+1, arr[len(arr)-1-i])
	}
}
