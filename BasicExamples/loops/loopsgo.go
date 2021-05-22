package main

import "fmt"

func main() {
	i := 0

	for i < 5 {
		fmt.Print(i)
		i++
	}

	fmt.Println("")

	for a := 5; a >= 0; a-- {
		for b := a; b >= 0; b-- {
			fmt.Print("*")
		}
		fmt.Println("")
	}

}
