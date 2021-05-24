package main

import (
	"fmt"
	"math"
)

func main() {

	arr := []int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90}
	var ele int
	fmt.Scanf("%d", &ele)
	fmt.Println("")
	var found bool
	var index int
	var low, high, mid int = 0, len(arr) - 1, 0
	//fmt.Printf("ele=%v,low=%v,high=%v,mid=%v", ele, low, high, mid)

	for low <= high {
		fmt.Printf("ele=%v,low=%v,high=%v,mid=%v\n", ele, low, high, mid)
		mid = int(math.Ceil((float64)(low+high) / 2))
		if ele == arr[mid] {
			found = true
			index = mid
			break
		} else if ele > arr[mid] {
			low = mid
		} else {
			high = mid
		}
	}
	if found {
		fmt.Printf("The element %v is found at the index %v", ele, index)
	} else {
		fmt.Println("The element is not present in the array")
	}
}
