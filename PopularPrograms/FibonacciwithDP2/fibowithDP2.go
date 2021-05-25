package main

import "fmt"

func fibonacci(n int, arr []int) int {
	if n == 0 || n == 1 {
		return n
	}
	if arr[n] == 0 {
		arr[n] = fibonacci(n-1, arr) + fibonacci(n-2, arr)
	}
	fmt.Printf("%02d\n", arr)
	return arr[n]
}

func fibonacciwithdp(n int) int {
	arr := make([]int, n+1)
	return fibonacci(n, arr)
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	fmt.Println(fibonacciwithdp(n))
}
