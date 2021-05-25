package main

import "fmt"

func findFibonacci(n int, arr []int) int {

	arr[0] = 0

	arr[1] = 1

	for i := 2; i < n; i++ {
		//fmt.Println(arr[i], i)
		arr[i] = arr[i-1] + arr[i-2]
		fmt.Println("values=", arr[i], i)

		//fmt.Println(arr[k-1])
	}
	return (arr[n-1] + arr[n-2])
}

func main() {

	var k int
	fmt.Scanf("%d", &k)
	arr := make([]int, k)
	println(findFibonacci(k, arr))

}
