package main

import "fmt"

func main() {

	var mat1 [3][3]int
	var mat2 [3][3]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scanf("%d", &mat1[i][j])
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scanf("%d", &mat2[i][j])
		}
	}
	fmt.Println(mat1)
	fmt.Println(mat2)
	fmt.Println("the sum of the matrix will be")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(mat1[i][j] + mat2[i][j])
			fmt.Print(" ")

		}
		fmt.Println("")
	}
	x := 5
	y := 10
	fmt.Println("Value of x & y = ", x, y)
	x, y = y, x
	fmt.Println(x, y)
}
