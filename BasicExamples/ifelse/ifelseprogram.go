package main

import "fmt"

func main() {
	var i int
	fmt.Scanf("%d", &i)

	if i%2 == 0 {
		fmt.Println("i is even")
	} else {
		fmt.Println("i is odd")
	}
}
